package electionRole

import (
	"PTIT_TN/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.ElectionRole, error) {
	var electionRoles []*entities.ElectionRole
	query := i.db.Gdb().
		WithContext(c).Model(&entities.ElectionRole{})
	err := query.Find(&electionRoles).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.ElectionRole{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Role Repo] Find All Error: %v", err))
		return nil, err
	}
	return electionRoles, nil
}

func (i impl) FindById(c *gin.Context, id uint64) (*entities.ElectionRole, error) {
	electionRole := &entities.ElectionRole{}
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionRole{}).
		Where("id = ?", id)
	err := query.First(electionRole).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Role Repo] Find By Id Error: %v", err))
		return nil, err
	}
	return electionRole, nil
}

func (i impl) FindByOption(c *gin.Context, option entities.ElectionRole) ([]*entities.ElectionRole, error) {
	var electionRoles []*entities.ElectionRole
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionRole{}).
		Where(&option)
	err := query.Find(&electionRoles).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.ElectionRole{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Role Repo] Find By Option Error: %v", err))
		return nil, err
	}
	return electionRoles, nil
}

func (i impl) FindByElectionCandidateId(c *gin.Context, electionCandidateId uint64) ([]*entities.ElectionRole, error) {
	var electionRoles []*entities.ElectionRole
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionRole{}).
		Where("election_candidate_id = ?", electionCandidateId)
	err := query.Find(&electionRoles).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.ElectionRole{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Role Repo] Find By Election Candidate Id Error: %v", err))
		return nil, err
	}
	return electionRoles, nil
}

func (i impl) FindByElectionId(c *gin.Context, electionId uint64) ([]*entities.ElectionRole, error) {
	var electionRoles []*entities.ElectionRole
	tx := i.db.Gdb().WithContext(c).Raw("select er.* from ptit_tn.election_roles er "+
		"left join ptit_tn.election_candidates ec on er.election_candidate_id = ec.id "+
		"left join ptit_tn.elections e on e.id = ec.election_id "+
		"where e.id=?;", electionId)
	err := tx.Scan(&electionRoles).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.ElectionRole{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Role Repo] Find By Election Id Error: %v", err))
		return nil, err
	}
	return electionRoles, nil
}
