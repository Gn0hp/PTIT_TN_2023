package entities

type ElectionRole struct {
	DefaultModel
	Name                string `json:"name"`
	ElectionCandidateId uint64 `json:"election_candidate_id"`
	Note                string `json:"note"`
}

type ElectionRoleDto struct {
	ElectionRole
}

func (e *ElectionRoleDto) ToEntity() (interface{}, error) {
	return &e.ElectionRole, nil
}
func (e *ElectionRoleDto) IsValid() bool {
	return true
}

func IsAllElectionRoleValid(rs []ElectionRole) bool {
	for _, role := range rs {
		tmp := ElectionRoleDto{ElectionRole: role}
		if !tmp.IsValid() {
			return false
		}

	}
	return true
}

func ElectionRoleToEntitiesWithEcId(rs []ElectionRole, id uint64) ([]*ElectionRole, error) {
	var entities []*ElectionRole
	for _, role := range rs {
		role.ElectionCandidateId = id
		tmp := ElectionRoleDto{ElectionRole: role}
		entity, err := tmp.ToEntity()
		if err != nil {
			return nil, err
		}
		e := entity.(*ElectionRole)
		entities = append(entities, e)
	}
	return entities, nil
}
