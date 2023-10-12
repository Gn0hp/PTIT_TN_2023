package entities

type Election struct {
	DefaultModel
	ApiObjectId uint64 `json:"api_object_id"` // id of election on blockchain
}
