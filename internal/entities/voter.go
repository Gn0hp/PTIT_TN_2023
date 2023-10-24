package entities

import "PTIT_TN/pkg/utils"

type Voter struct {
	User
}

type VoterDto struct {
	Voter
}

func (u *VoterDto) ToEntity() (interface{}, error) {
	return &u.Voter, nil
}

func (u *VoterDto) IsValid() bool {
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
