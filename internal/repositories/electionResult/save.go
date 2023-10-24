package electionResult

import (
	"PTIT_TN/internal/entities"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(c *gin.Context, v *entities.ElectionResult) error {
	return i.db.Gdb().WithContext(c).Model(&entities.ElectionResult{}).Create(v).Error
}
func (i impl) UpdateById(c *gin.Context, v *entities.ElectionResult) error {
	return i.db.Gdb().WithContext(c).Model(&entities.ElectionResult{}).Where("id = ?", v.ID).Updates(v).Error
}
