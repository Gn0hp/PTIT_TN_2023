package admin

import (
	"PTIT_TN/pkg"
	"PTIT_TN/pkg/rabbitMQ"
)

func GetNonceBlockchainService(h *Handler) interface{} {
	resultChan, errChan := make(chan rabbitMQ.RabbitRequest), make(chan error)
	h.rabbitMQ.Send(rabbitMQ.RabbitRequest{
		Type: pkg.TxTypeNonce,
		Data: nil,
	})
	data := h.rabbitMQ.Receive(resultChan, errChan)
	return data
}

//
//// electionId
//func GetElectionResultService(h *Handler, electionId int64) interface{} {
//	resultChan, errChan := make(chan rabbitMQ.RabbitRequest), make(chan error)
//	h.rabbitMQ.Send(rabbitMQ.RabbitRequest{
//		Type: pkg.TxTypeGetElectionResult,
//		Data: map[string]interface{}{
//			"election_id": electionId,
//		},
//	})
//	go h.rabbitMQ.Receive(resultChan, errChan)
//	select {
//	case res := <-resultChan:
//		return res
//	case err := <-errChan:
//		h.logger.Error(fmt.Sprintf("[Election Handler] Get Election Result failed, detail: %v", err))
//		return nil
//	}
//}
//
//// electionId - election_candidate_id
//func GetElectionToResultService(h *Handler, electionId, candidateId int64) interface{} {
//	resultChan, errChan := make(chan rabbitMQ.RabbitRequest), make(chan error)
//	h.rabbitMQ.Send(rabbitMQ.RabbitRequest{
//		Type: pkg.TxTypeElectionToResult,
//		Data: map[string]interface{}{
//			"election_id":  electionId,
//			"candidate_id": candidateId,
//		},
//	})
//	go h.rabbitMQ.Receive(resultChan, errChan)
//	select {
//	case res := <-resultChan:
//		return res
//	case err := <-errChan:
//		h.logger.Error(fmt.Sprintf("[Election Handler] Get Election Result failed, detail: %v", err))
//		return nil
//	}
//}
