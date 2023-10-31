package entities

type BallotRoleElect struct {
	DefaultModel
	BallotId    uint64 `json:"ballot_id"`
	RoleElectId uint64 `json:"role_elect_id"`
	Note        string `json:"note"`
}

type BallotRoleElectDto struct {
	BallotRoleElect
}

func (e *BallotRoleElectDto) ToEntity() (interface{}, error) {
	return &e.BallotRoleElect, nil
}
