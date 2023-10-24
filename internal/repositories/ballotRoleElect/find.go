package ballotRoleElect

import (
	"PTIT_TN/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.BallotRoleElect, error) {
	var result []*entities.BallotRoleElect
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.BallotRoleElect{})
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Find BallotRoleElect failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindById(c *gin.Context, id uint64) (*entities.BallotRoleElect, error) {
	result := &entities.BallotRoleElect{}
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.BallotRoleElect{}).
		Where("id = ?", id)
	err := query.First(result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Find BallotRoleElect failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindByOption(c *gin.Context, option entities.BallotRoleElect) ([]*entities.BallotRoleElect, error) {
	var result []*entities.BallotRoleElect
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.BallotRoleElect{}).
		Where(&option)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Find BallotRoleElect failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindByBallotId(c *gin.Context, ballotId uint64) ([]*entities.BallotRoleElect, error) {
	var result []*entities.BallotRoleElect
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.BallotRoleElect{}).
		Where("ballot_id = ?", ballotId)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Find BallotRoleElect failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindByRoleElectId(c *gin.Context, roleElectId uint64) ([]*entities.BallotRoleElect, error) {
	var result []*entities.BallotRoleElect
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.BallotRoleElect{}).
		Where("role_elect_id = ?", roleElectId)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Find BallotRoleElect failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
