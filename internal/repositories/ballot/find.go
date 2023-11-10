package ballot

import (
	"PTIT_TN/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.Ballot, error) {
	var result []*entities.Ballot
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Ballot{})
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Ballot Repo] Find Ballot failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindById(c *gin.Context, id uint64) (*entities.Ballot, error) {
	ballot := &entities.Ballot{}
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Ballot{}).
		Where("id = ?", id)
	err := query.First(ballot).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Ballot Repo] Find Ballot failed, detail: %v", err))
		return nil, err
	}
	return ballot, nil
}
func (i impl) FindByOption(c *gin.Context, option entities.Ballot) ([]*entities.Ballot, error) {
	var result []*entities.Ballot
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Ballot{}).
		Where(&option)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Ballot Repo] Find Ballot failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindByVoterId(c *gin.Context, voterId uint64) (*entities.Ballot, error) {
	var result *entities.Ballot
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Ballot{}).
		Where("voter_id = ?", voterId)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Ballot Repo] Find Ballot failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindByResultId(c *gin.Context, resultId uint64) ([]*entities.Ballot, error) {
	var result []*entities.Ballot
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Ballot{}).
		Where("result_id = ?", resultId)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Ballot Repo] Find Ballot failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
