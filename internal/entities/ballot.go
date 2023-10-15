package entities

type Ballot struct {
	DefaultModel

	VoterId uint64 `json:"voter_id"`
	CandidateId uint64 `json:"candidate_id"`
}

type BallotDto struct {
	Ballot
}