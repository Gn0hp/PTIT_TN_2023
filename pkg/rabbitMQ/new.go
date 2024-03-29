package rabbitMQ

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"log"
	"logur.dev/logur"
	"time"
)

type RabbitRequest struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
type MqService interface {
	GetConnection() *amqp091.Connection
	Send(body RabbitRequest)
	Receive(result chan<- RabbitRequest, errChan chan error) *RabbitRequest
}
type impl struct {
	logger logur.LoggerFacade
	ctx    context.Context
	conn   *amqp091.Connection
}

func NewConnection(c Config, logger logur.LoggerFacade, ctx context.Context) (MqService, error) {
	conn, err := amqp091.Dial(c.ConnectionString())
	if err != nil {
		return nil, err
	}
	logger.Info("rabbitMq connected !!")
	return &impl{
		logger: logger,
		conn:   conn,
		ctx:    ctx,
	}, nil
}
func (i impl) Send(body RabbitRequest) {
	ch, err := i.conn.Channel()
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Rabbit MQ Server] Error sending to service, detail: %v", err))
		return
	}
	q, err := ch.QueueDeclare("rabbit_ptit_tn_prj", false, false, false, false, nil)
	if err != nil {
		logrus.Errorf("")
	}
	bodyBytes, _ := json.Marshal(body)
	err = ch.PublishWithContext(i.ctx, "", q.Name, false, false, amqp091.Publishing{
		Headers:         nil,
		ContentType:     "text/json",
		ContentEncoding: "",
		DeliveryMode:    0,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "",
		Timestamp:       time.Time{},
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            bodyBytes,
	})
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Rabbit MQ Server] Error sending to service, detail: %v", err))
		return
	}
	i.logger.Info("[Rabbit MQ] Send success")

}
func (i impl) Receive(result chan<- RabbitRequest, errChan chan error) *RabbitRequest {
	ch, err := i.conn.Channel()
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Rabbit MQ Server] Error sending to service, detail: %v", err))
		return nil
	}
	q, err := ch.QueueDeclare("rabbit_ptit_tn_prj_reverse", false, false, false, false, nil)
	if err != nil {
		logrus.Errorf("")
	}
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Rabbit MQ Server] Error sending to service, detail: %v", err))
		return nil
	}
	for d := range msgs {
		var body RabbitRequest
		log.Printf("Received a message from blockchain: %s", d.Body)
		err := json.Unmarshal(d.Body, &body)
		if err != nil {
			i.logger.Error(fmt.Sprintf("[Rabbit MQ Server] Error sending to service, detail: %v", err))
			errChan <- err
			continue
		}
		if body.Data != nil {
			return &body
		}
	}
	return nil
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.User, c.Password, c.Host, c.Port)
}

func (i impl) GetConnection() *amqp091.Connection {
	return i.conn
}
