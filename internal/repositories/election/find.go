package election

import (
	"PTIT_TN/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.Election, error) {
	var elections []*entities.Election
	query := i.db.Gdb().
		WithContext(c).Model(&entities.Election{})
	err := query.Find(&elections).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Election{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Repo] Find All Election failed, detail: %v", err))
		return nil, err
	}
	return elections, nil
}

func (i impl) FindById(c *gin.Context, id uint64) (*entities.Election, error) {
	elections := &entities.Election{}
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Election{}).
		Where("id = ?", id)
	err := query.First(elections).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Repo] Find Election failed, detail: %v", err))
		return nil, err
	}
	return elections, nil
}

func (i impl) FindByOption(c *gin.Context, option entities.Election) ([]*entities.Election, error) {
	var elections []*entities.Election
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Election{}).
		Where(&option)
	err := query.Find(&elections).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Election{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Repo] Find Election failed, detail: %v", err))
		return nil, err
	}
	return elections, nil
}

func (i impl) FindOpenElection(c *gin.Context) (*entities.Election, error) {
	var election entities.Election
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Election{}).
		Where("status = ?", entities.ELECTION_STATUS_REGISTERING).
		Or("status = ?", entities.ELECTION_STATUS_OPENING)
	err := query.First(&election).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Repo] Find Open Election failed, detail: %v", err))
		return nil, err
	}
	return &election, nil
}
