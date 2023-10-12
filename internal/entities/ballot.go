package entities

type Ballot struct {
	DefaultModel

	VoterId uint64 `json:"voter_id"`
}
