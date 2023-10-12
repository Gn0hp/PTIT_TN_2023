package entities

type Admin struct {
	DefaultModel
	Username string `json:"username,omitempty" gorm:"size:32; DEFAULT NULL"`
	Password string `json:"password" gorm:"size:256; NOT NULL"`
	Status   string `json:"status" gorm:"size:16; NOT NULL"`
}
