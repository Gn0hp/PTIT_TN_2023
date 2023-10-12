package entities

import (
	"PTIT_TN/pkg/utils"
	"time"
)

type User struct {
	DefaultModel
	FullName    string    `json:"full_name" gorm:"size:128; NOT NULL"`
	Gender      string    `json:"gender" gorm:"size:16; NOT NULL"`
	DateOfBirth time.Time `json:"date_of_birth" gorm:"size:32; NOT NULL"`
	Address     string    `json:"address" gorm:"size:64; NOT NULL"`
	CccdID      string    `json:"cccd_id" gorm:"size:32; NOT NULL; uniqueIndex"`
	CccdDate    time.Time `json:"cccd_date" gorm:"size:32; NOT NULL"`
	CccdPlace   string    `json:"cccd_place" gorm:"size:64; NOT NULL"`
	Email       string    `json:"email" gorm:"size:128; NOT NULL"`
	Phone       string    `json:"phone" gorm:"size:64; NOT NULL"`
	Password    string    `json:"password" gorm:"size:256; NOT NULL"`
	Status      string    `json:"status" gorm:"size:16; NOT NULL"`
	Note        string    `json:"note,omitempty" gorm:"size:512; default:NULL"`
}

type UserDto struct {
	User
}

func (u *UserDto) ToEntity() (interface{}, error) {
	return &u.User, nil
}

func (u *UserDto) IsValid() bool {
	if u.FullName == "" || u.Gender == "" || u.Address == "" || u.CccdID == "" || u.CccdPlace == "" {
		return false
	}
	if u.CccdDate.IsZero() || u.DateOfBirth.IsZero() {
		return false
	}
	if u.Email == "" || u.Phone == "" || u.Password == "" || !utils.ValidateStringNumber(u.Phone) || !utils.ValidateMail(u.Email) {
		return false
	}
	return true
}
