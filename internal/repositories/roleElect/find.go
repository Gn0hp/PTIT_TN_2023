package roleElect

import (
	"PTIT_TN/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.RoleElect, error) {
	var roleElects []*entities.RoleElect
	query := i.db.Gdb().
		WithContext(c).Model(&entities.RoleElect{})
	err := query.Find(&roleElects).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.RoleElect{}, nil
		}
		i.logger.Error(fmt.Sprintf("[RoleElect Repo] Find All RoleElect failed, detail: %v", err))
		return nil, err
	}
	return roleElects, nil
}

func (i impl) FindById(c *gin.Context, id uint64) (*entities.RoleElect, error) {
	roleElect := &entities.RoleElect{}
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.RoleElect{}).
		Where("id = ?", id)
	err := query.First(roleElect).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[RoleElect Repo] Find RoleElect failed, detail: %v", err))
		return nil, err
	}
	return roleElect, nil
}

func (i impl) FindByCandidateId(c *gin.Context, candidateId uint64) ([]*entities.RoleElect, error) {
	var roleElects []*entities.RoleElect
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.RoleElect{}).
		Where("candidate_id = ?", candidateId)
	err := query.Find(&roleElects).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.RoleElect{}, nil
		}
		i.logger.Error(fmt.Sprintf("[RoleElect Repo] Find RoleElect failed, detail: %v", err))
		return nil, err
	}
	return roleElects, nil
}

func (i impl) FindByElectionRoleId(c *gin.Context, electionRoleId uint64) ([]*entities.RoleElect, error) {
	var roleElects []*entities.RoleElect
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.RoleElect{}).
		Where("election_role_id = ?", electionRoleId)
	err := query.Find(&roleElects).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.RoleElect{}, nil
		}
		i.logger.Error(fmt.Sprintf("[RoleElect Repo] Find RoleElect failed, detail: %v", err))
		return nil, err
	}
	return roleElects, nil
}

func (i impl) FindByOption(c *gin.Context, option entities.RoleElect) ([]*entities.RoleElect, error) {
	var roleElects []*entities.RoleElect
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.RoleElect{}).
		Where(&option)
	err := query.Find(&roleElects).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.RoleElect{}, nil
		}
		i.logger.Error(fmt.Sprintf("[RoleElect Repo] Find RoleElect failed, detail: %v", err))
		return nil, err
	}
	return roleElects, nil
}
