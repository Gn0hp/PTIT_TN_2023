package entities

const (
	CANDIDATE_STATUS_PENDING  = "PENDING"  // chờ duyệt
	CANDIDATE_STATUS_VERIFIED = "VERIFIED" // đã duyệt
	CANDIDATE_STATUS_REJECTED = "REJECTED" // bị từ chối
	CANDIDATE_STATUS_BLOCKED  = "BLOCKED"  // bị khóa
	CANDIDATE_STATUS_ACTIVE   = "ACTIVE"   // đang ứng cử
	CANDIDATE_STATUS_INACTIVE = "INACTIVE" // không ứng cử

)

type Candidate struct {
	User

	CccdID          string `json:"cccd_id" gorm:"size:32; NOT NULL;"`
	CandidateStatus string `json:"candidate_status" gorm:"size:16; NOT NULL"` // ["active", "inactive", "rejected", blocked]
	Party           string `json:"party"`
	Patrimony       string `json:"patrimony,omitempty" gorm:"size:512; default:NULL"` // for Candidate only
}

type CandidateDto struct {
	Candidate
}

func (c *CandidateDto) ToEntity() (interface{}, error) {
	return &c.Candidate, nil
}
func (c *CandidateDto) IsValid() bool {
	return true
}
