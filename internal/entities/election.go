package entities

type Election struct {
	DefaultModel
	ApiObjectId uint64 `json:"api_object_id"` // id of election on blockchain

	DateEndRegister int64 `json:"date_end_register"`
	DateStartElecting int64 `json:"date_start_electing"`
	Duration int64 `json:"duration"`

	Status int64 `json:"status"`
	Note string `json:"note"`
}

type ElectionDto struct {
	Election
}