package election

import (
	"PTIT_TN/internal/entities"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(c *gin.Context, e *entities.Election) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Election{}).Create(e).Error
}

func (i impl) UpdateById(c *gin.Context, e *entities.Election) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Election{}).Where("id = ?", e.ID).Updates(e).Error
}
