package electionCandidate

import (
	"PTIT_TN/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.ElectionCandidate, error) {
	var result []*entities.ElectionCandidate
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionCandidate{})
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.ElectionCandidate{}, nil
		}
		i.logger.Error(fmt.Sprintf("[ElectionCandidate Repo] FindAll Error: %s", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindById(c *gin.Context, id uint64) (*entities.ElectionCandidate, error) {
	result := &entities.ElectionCandidate{}
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionCandidate{}).
		Where("id = ?", id)
	err := query.First(result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[ElectionCandidate Repo] FindById Error: %s", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindByElectionId(c *gin.Context, electionId uint64) (*entities.ElectionCandidate, error) {
	var result *entities.ElectionCandidate
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionCandidate{}).
		Where("election_id = ?", electionId)
	err := query.First(&result).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[ElectionCandidate Repo] FindByElectionId Error: %s", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindByOption(c *gin.Context, option entities.ElectionCandidate) ([]*entities.ElectionCandidate, error) {
	var result []*entities.ElectionCandidate
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionCandidate{}).
		Where(&option)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.ElectionCandidate{}, nil
		}
		i.logger.Error(fmt.Sprintf("[ElectionCandidate Repo] FindByOption Error: %s", err))
		return nil, err
	}
	return result, nil
}
