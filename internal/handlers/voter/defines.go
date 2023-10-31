package voter

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/repositories/candidate"
)

type ViewCandidateResponse struct {
	MapCandidate map[string][]*candidate.ViewCandidateDbFind `json:"map_candidate"`
	SupportedMap map[string]*entities.ElectionRole           `json:"supported_map"`
}

type RegisterCandidateRequest struct {
	//Candidate entities.Candidate
	Role uint64 `json:"role"`
	entities.User
	Patrimony string `json:"patrimony"`
	Party     string `json:"party"`
}

type VoteRequestBody struct {
	ElectionId uint64        `json:"election_id"`
	VoterId    uint64        `json:"voter_id"`
	RoleElects []RoleElectRB `json:"role_elects"`
}
type RoleElectRB struct {
	CandidateId uint64 `json:"candidate_id"`
	RoleElectId uint64 `json:"role_elect_id"`
}
