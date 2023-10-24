package entities

type RoleElect struct {
	DefaultModel
	Name                string `json:"name"`
	CandidateId         uint64 `json:"candidate_id"`
	ElectionCandidateId uint64 `json:"election_candidate_id"`
	Note                string `json:"note"`
}
