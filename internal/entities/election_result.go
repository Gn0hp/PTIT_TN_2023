package entities

type ElectionResult struct {
	DefaultModel
	ElectionCandidateId uint64 `json:"election_candidate_id"`
	Note                string `json:"note"`
}

type ElectionResultDto struct {
	ElectionResult
}
