package entities

type ElectionResult struct {
	DefaultModel

	ElectionCandidateId uint64 `json:"election_candidate_id"`
	BallotId            uint64 `json:"ballot_id"`
}

type ElectionResultDto struct {
	ElectionResult
}