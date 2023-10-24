package post

import (
	"PTIT_TN/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.Post, error) {
	var result []*entities.Post
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Post{})
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Post{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Post Repo] Find All Post failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindById(c *gin.Context, id uint64) (*entities.Post, error) {
	post := &entities.Post{}
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Post{}).
		Where("id = ?", id)
	err := query.First(post).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Post Repo] Find Post failed, detail: %v", err))
		return nil, err
	}
	return post, nil
}
func (i impl) FindByCandidateId(c *gin.Context, candidateId uint64) ([]*entities.Post, error) {
	var result []*entities.Post
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Post{}).
		Where("candidate_id = ?", candidateId)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Post{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Post Repo] Find Post failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindByOption(c *gin.Context, option entities.Post) ([]*entities.Post, error) {
	var result []*entities.Post
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Post{}).
		Where(&option)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Post{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Post Repo] Find Post failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
