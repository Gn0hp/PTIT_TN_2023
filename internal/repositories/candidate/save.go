package candidate

import (
	"PTIT_TN/internal/entities"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(c *gin.Context, v *entities.Candidate) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Candidate{}).Create(v).Error
}

func (i impl) UpdateById(c *gin.Context, id uint64, v *entities.Candidate) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Candidate{}).Where("id = ?", id).Updates(v).Error
}

func (i impl) SelfUpdate(c *gin.Context, v *entities.Candidate) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Candidate{}).Where("id = ?", v.ID).Updates(v).Error
}
func (i impl) Delete(c *gin.Context, id uint64) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Candidate{}).Where("id = ?", id).Delete(&entities.Candidate{}).Error
}
