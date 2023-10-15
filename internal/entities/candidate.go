package entities

type Candidate struct {
	User

	RoleElect string `json:"role_elect" gorm:"gorm:size:32"`
	Party     string
	Patrimony string `json:"patrimony,omitempty" gorm:"size:512; default:NULL"` // for Candidate only
}

type CandidateDto struct {
	Candidate
}