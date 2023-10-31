package entities

const (
	ELECTION_STATUS_REGISTERING = "REGISTERING"
	ELECTION_STATUS_OPENING     = "OPENING"
	ELECTION_STATUS_FINISHED    = "FINISHED"
)

type Election struct {
	DefaultModel
	ApiObjectId uint64 `json:"api_object_id"` // id of election on blockchain

	Organization      string `json:"organization"`
	DateEndRegister   int64  `json:"date_end_register"`
	DateStartElecting int64  `json:"date_start_electing"`
	Duration          int64  `json:"duration"`
	Status            string `json:"status"` // ['REGISTERING', 'OPENING', 'FINISHED']
	Note              string `json:"note"`
}

type ElectionDto struct {
	Election
}

func (e *ElectionDto) IsValid() bool {
	return true
}
func (e *ElectionDto) ToEntity() (interface{}, error) {
	return &e.Election, nil
}
