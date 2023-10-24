package voter

import (
	"PTIT_TN/internal/entities"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.Voter, error) {
	var voters []*entities.Voter
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Voter{})
	err := query.Find(&voters).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[User Repo] Find All Voter failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Voter{}, nil
		}
		return nil, err
	}
	return voters, nil
}
func (i impl) FindByCccd(c *gin.Context, cccd string) (*entities.Voter, error) {
	user := &entities.Voter{}

	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Voter{}).
		Where("cccd_id = ?", cccd)
	err := query.First(user).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[User Repo] Find User failed, detail: %v", err))
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("voter not found")
		}
		return nil, err
	}
	return user, nil

}
func (i impl) FindById(c *gin.Context, id uint64) (*entities.Voter, error) {
	user := &entities.Voter{}

	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Voter{}).
		Where("id = ?", id)
	err := query.First(user).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[User Repo] Find User failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("voter not found")
		}
		return nil, err
	}
	return user, nil
}

func (i impl) GetPendingVoters(c *gin.Context) ([]*entities.Voter, error) {
	var voters []*entities.Voter
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Voter{}).
		Where("status = ?", entities.REGISTER_STATUS_PENDING)
	err := query.Find(&voters).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[User Repo] Find Pending Voter failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Voter{}, nil
		}
		return nil, err
	}
	return voters, nil
}
