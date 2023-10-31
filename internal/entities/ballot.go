package entities

type Ballot struct {
	DefaultModel

	VoterId  uint64 `json:"voter_id"`
	ResultId uint64 `json:"result_id"`
	Note     string `json:"note"`
}

type BallotDto struct {
	Ballot
}

func (b *BallotDto) ToEntity() (interface{}, error) {
	return &b.Ballot, nil
}
