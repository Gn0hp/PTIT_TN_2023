package post

import (
	"PTIT_TN/internal/entities"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(c *gin.Context, v *entities.Post) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Post{}).Create(v).Error
}
func (i impl) UpdateById(c *gin.Context, v *entities.Post) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Post{}).Where("id = ?", v.ID).Updates(v).Error
}
