package election

import "PTIT_TN/internal/entities"

type CreateElectionBody struct {
	entities.Election
	Roles []entities.ElectionRole `json:"roles"`
}

type CreateElectionReq struct {
	ID           uint64 `json:"id"`
	StartDate    uint64 `json:"start_date"`
	Duration     uint64 `json:"duration"`
	NumCandidate uint64 `json:"num_candidate"`
}
