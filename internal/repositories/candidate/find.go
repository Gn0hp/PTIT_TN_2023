package candidate

import (
	"PTIT_TN/internal/entities"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindByCccd(c *gin.Context, cccd string) (*entities.Candidate, error) {
	user := &entities.Candidate{}

	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Candidate{}).
		Where("cccd_id = ?", cccd)
	err := query.First(user).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find Candidate failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("candidate not found")
		}
		return nil, err
	}
	return user, nil

}
func (i impl) FindById(c *gin.Context, id uint64) (*entities.Candidate, error) {
	user := &entities.Candidate{}

	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Candidate{}).
		Where("id = ?", id)
	err := query.First(user).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find Candidate failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("candidate not found")
		}
		return nil, err
	}
	return user, nil
}
func (i impl) FindAll(c *gin.Context) ([]*entities.Candidate, error) {
	var candidates []*entities.Candidate
	query := i.db.Gdb().
		WithContext(c).Model(&entities.Candidate{})
	err := query.Find(&candidates).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find All Candidate failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Candidate{}, nil
		}
		return nil, err
	}
	return candidates, nil
}
func (i impl) FindByOption(c *gin.Context, option entities.Candidate) ([]*entities.Candidate, error) {
	var candidates []*entities.Candidate
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Candidate{}).
		Where(&option)
	err := query.Find(&candidates).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find Candidate failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Candidate{}, nil
		}
		return nil, err
	}
	return candidates, nil
}
func (i impl) GetPendingCandidate(c *gin.Context) ([]*entities.Candidate, error) {
	var candidates []*entities.Candidate
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Candidate{}).
		Where("status = ?", entities.CANDIDATE_STATUS_PENDING)
	err := query.Find(&candidates).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[User Repo] Find Pending Candidate failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Candidate{}, nil
		}
		return nil, err
	}
	return candidates, nil
}
