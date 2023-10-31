package entities

type RoleElect struct {
	DefaultModel
	Name           string `json:"name"`
	CandidateId    uint64 `json:"candidate_id"`
	ElectionRoleId uint64 `json:"election_role_id"`
	Note           string `json:"note"`
}

type RoleElectDto struct {
	RoleElect
}

func (r *RoleElectDto) ToEntity() (interface{}, error) {
	return &r.RoleElect, nil
}
func (r *RoleElectDto) IsValid() bool {
	return true
}

func IsAllRoleElectValid(rs []RoleElect) bool {
	for _, role := range rs {
		tmp := RoleElectDto{RoleElect: role}
		if !tmp.IsValid() {
			return false
		}

	}
	return true
}

func RoleElectToEntities(rs []RoleElect) ([]*RoleElect, error) {
	var entities []*RoleElect
	for _, role := range rs {
		tmp := RoleElectDto{RoleElect: role}
		entity, err := tmp.ToEntity()
		if err != nil {
			return nil, err
		}
		e := entity.(*RoleElect)
		entities = append(entities, e)
	}
	return entities, nil
}
