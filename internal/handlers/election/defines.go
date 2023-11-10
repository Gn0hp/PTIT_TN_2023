package election

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/repositories/candidate"
)

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

type ViewCandidateResponse struct {
	MapCandidate map[string][]*candidate.ViewResultDbFind `json:"map_candidate"`
	SupportedMap map[string]*entities.ElectionRole        `json:"supported_map"`
}
