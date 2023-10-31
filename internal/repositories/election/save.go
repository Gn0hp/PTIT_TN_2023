package election

import (
	"PTIT_TN/internal/entities"
	"errors"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(c *gin.Context, e *entities.Election) error {
	rs := i.db.Gdb().WithContext(c).Model(entities.Election{}).Create(e)
	if rs.Error != nil {
		return rs.Error
	}
	if rs.RowsAffected == 0 {
		return errors.New("create election failed")
	}
	return nil
}

func (i impl) UpdateById(c *gin.Context, e *entities.Election) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Election{}).Where("id = ?", e.ID).Updates(e).Error
}
