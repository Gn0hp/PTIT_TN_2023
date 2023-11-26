package main

import (
	election_contract "PTIT_TN/contracts"
	log2 "PTIT_TN/pkg/logger"
	"PTIT_TN/pkg/rabbitMQ"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"os"
	"time"
)

type Request struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
type CreateElectionReq struct {
	ID           uint64 `json:"id"`
	StartDate    uint64 `json:"start_date"`
	Duration     uint64 `json:"duration"`
	NumCandidate uint64 `json:"num_candidate"`
}
type Ballot struct {
	ElectionId  uint64  `json:"election_id"`
	VoterId     string  `json:"voter_id"`
	ChoiceValue []int64 `json:"choice_value"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.TODO()
	client, err := ethclient.Dial(os.Getenv("NODE_RPC_URL_BSC"))
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to the Binance Smart Chain network: %v", err))
	}
	logger := log2.NewLogger(log2.Config{
		Format:        "logfmt",
		Level:         "info",
		NoColor:       false,
		FullTimestamp: true,
	})
	chainID, accountBound, err := initClient(ctx, client)
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := election_contract.NewElectionContract(contractAddress, client)
	accountBoundRead := &bind.CallOpts{
		From: common.HexToAddress(os.Getenv("PUBLIC_KEY")),
	}
	contract := New(logger, chainID, accountBound, accountBoundRead, instance)
	fmt.Printf("Contract is loaded: %v\n", contract)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to instantiate a smart contract: %v", err))
	}
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}
	q, err := ch.QueueDeclare("rabbit_ptit_tn_prj", false, false, false, false, nil)
	if err != nil {
		log.Panic(err)
	}
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
	}
	forever := make(chan bool)
	var jsonRes Request
	go func() {
		for d := range msgs {
			json.Unmarshal(d.Body, &jsonRes)
			log.Printf("Received a message: %s", d.Body)
			switch jsonRes.Type {
			case TxTypeCreateElection:
				fmt.Println("Create election ")
				createElectionReq, ok := jsonRes.Data.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid json object")
				}
				data := CreateElectionReq{
					ID:           uint64(createElectionReq["id"].(float64)),
					StartDate:    uint64(createElectionReq["start_date"].(float64)),
					Duration:     uint64(createElectionReq["duration"].(float64)),
					NumCandidate: uint64(createElectionReq["num_candidate"].(float64)),
				}
				tx, err := contract.CreateNewElection(data.StartDate, data.Duration, data.NumCandidate)
				if err != nil {
					logger.Error(fmt.Sprintf("[Contract Instance] Create New Election failed, detail: %v", err))
					continue
				}
				logger.Info(fmt.Sprintf("[Contract Instance] Create New Election success, detail: %v", tx))
			case TxTypeVote:
				fmt.Printf("Vote: %v", jsonRes.Data)
				balJson, ok := jsonRes.Data.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid json object")
				}
				tmp, ok := balJson["choice_value"].([]interface{})
				if !ok {
					fmt.Println("Invalid json object: choice value")
				}
				var choiceValue []int64
				for _, v := range tmp {
					choiceValue = append(choiceValue, int64(v.(float64)))
				}
				ballot := Ballot{
					ElectionId:  uint64(balJson["election_id"].(float64)),
					VoterId:     balJson["voter_id"].(string),
					ChoiceValue: choiceValue,
				}
				res, err := contract.Vote(int64(ballot.ElectionId), ballot.VoterId, ballot.ChoiceValue)
				if err != nil {
					logger.Error(fmt.Sprintf("[Contract Instance] Vote failed, detail: %v", err))
					continue
				}
				logger.Info(fmt.Sprintf("[Contract Instance] Vote success, detail: %v", res))
			case TxTypeGetElectionResult:
				electionIdJson, ok := jsonRes.Data.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid json object")
				}
				electionId := int64(electionIdJson["election_id"].(float64))
				res, err := contract.GetElectionResult(electionId)
				if err != nil {
					logger.Error(fmt.Sprintf("[Contract Instance] Get Election Result failed, detail: %v", err))
					continue
				}
				Send(ch, rabbitMQ.RabbitRequest{
					Type: TxTypeGetElectionResult,
					Data: res,
				}, ctx)
				logger.Info(fmt.Sprintf("[Contract Instance] Get Election Result success, detail: %v", res))
			case TxTypeElectionToResult:
				dataJson, ok := jsonRes.Data.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid json object")
				}
				electionId := int64(dataJson["election_id"].(float64))
				choiceValue := int64(dataJson["candidate_id"].(float64))
				res, err := contract.ElectionToResult(electionId, choiceValue)
				if err != nil {
					logger.Error(fmt.Sprintf("[Contract Instance] Get Election Result failed, detail: %v", err))
					continue
				}
				Send(ch, rabbitMQ.RabbitRequest{
					Type: TxTypeElectionToResult,
					Data: res,
				}, ctx)
				logger.Info(fmt.Sprintf("[Contract Instance] Get Election Result success, detail: %v", res))
			case TxTypeTransferOwnership:
			case TxTypeRenounceOwnership:
			case TxTypeElectionMapping:
				res, err := contract.ElectionsMapping(1)
				if err != nil {
					fmt.Printf("Error: %v", err)
				}
				fmt.Printf(fmt.Sprintf("Data Election Mapping: %v", res))
			case TxTypeIsVoted:
			case TxTypeOwner:
				fmt.Printf("Data Owner: %v", contract.Owner())
			case TxTypeNonce:
				nonce, err := contract.GetNonce()
				if err != nil {
					logger.Error(fmt.Sprintf("[Contract Instance] Get Nonce failed, detail: %v", err))
					continue
				}
				Send(ch, rabbitMQ.RabbitRequest{
					Type: TxTypeNonce,
					Data: nonce.String(),
				}, ctx)
				logger.Info(fmt.Sprintf("[Contract Instance] Get Nonce success, detail: %v", nonce))
			case TxTypeSignature:
				signature, err := contract.GetSignatureString()
				if err != nil {
					logger.Error(fmt.Sprintf("[Contract Instance] Get Signature failed, detail: %v", err))
					continue
				}
				logger.Info(fmt.Sprintf("[Contract Instance] Get Signature success, detail: %v", signature))
			default:

			}
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func initClient(ctx context.Context, client *ethclient.Client) (*big.Int, *bind.TransactOpts, error) {
	prvKey := os.Getenv("PRIVATE_KEY")
	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to get chain id: %v", err))
		return nil, nil, err
	}
	accountBinding, err := TransactOptsAccountBinding(chainId, prvKey)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create account binding: %v", err))
		return nil, nil, err
	}
	return chainId, accountBinding, nil
}

func Send(ch *amqp091.Channel, body rabbitMQ.RabbitRequest, ctx context.Context) {
	q, err := ch.QueueDeclare("rabbit_ptit_tn_prj_reverse", false, false, false, false, nil)
	if err != nil {
		logrus.Errorf("")
	}
	bodyBytes, _ := json.Marshal(body)
	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp091.Publishing{
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
		logrus.Error(fmt.Sprintf("[Rabbit MQ Server] Error sending to service, detail: %v", err))
		return
	}
	logrus.Info("[Rabbit MQ] Send success")
}
