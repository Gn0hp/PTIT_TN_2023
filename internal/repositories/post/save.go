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

func (i impl) Delete(c *gin.Context, id uint64) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Post{}).Where("id = ?", id).Delete(&entities.Post{}).Error
}

func (i impl) Update(c *gin.Context, id uint64, v *entities.Post) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Post{}).Where("id = ?", id).Updates(v).Error
}
