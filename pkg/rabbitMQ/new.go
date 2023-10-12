package rabbitMQ

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"logur.dev/logur"
	"time"
)

type RabbitRequest struct{}
type MqService interface {
	GetConnection() *amqp091.Connection
	Send(ch *amqp091.Channel, body RabbitRequest)
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
func (i impl) Send(ch *amqp091.Channel, body RabbitRequest) {
	q, err := ch.QueueDeclare("rabbit_test_q", false, false, false, false, nil)
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

func (c Config) ConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.User, c.Password, c.Host, c.Port)
}

func (i impl) GetConnection() *amqp091.Connection {
	return i.conn
}
