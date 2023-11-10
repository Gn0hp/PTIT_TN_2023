package candidate

import "PTIT_TN/internal/entities"

type ViewCandidateDbFind struct {
	entities.Candidate
	RoleElectId    uint64 `json:"reId"`
	ElectionRoleId string `json:"erId"`
}
type ViewResultDbFind struct {
	entities.Candidate
	RoleElectId    uint64 `json:"reId"`
	ElectionRoleId string `json:"erId"`
	NumVote        uint64 `json:"num_vote"`
}
