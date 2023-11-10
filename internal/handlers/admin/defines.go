package admin

import (
	"PTIT_TN/internal/entities"
)

type StatByCandidateResponse struct {
	entities.Candidate    `json:"candidate"`
	entities.ElectionRole `json:"election_role"`
	TotalVote             int64  `json:"total_vote"`
	Role                  string `json:"role"`
	RoleElectId           uint64 `json:"role_elect_id"`
}
