package entities

type ElectionCandidate struct {
	DefaultModel

	ElectionId uint64 `json:"election_id" gorm:"type:INT(11);NOT NULL"`
	Note       string `json:"note"`
}

type ElectionCandidateDto struct {
	ElectionCandidate
}

func (e *ElectionCandidateDto) ToEntity() (interface{}, error) {
	return &e.ElectionCandidate, nil
}
