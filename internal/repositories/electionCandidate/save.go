package electionCandidate

import (
	"PTIT_TN/internal/entities"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(c *gin.Context, v *entities.ElectionCandidate) error {
	return i.db.Gdb().WithContext(c).Model(&entities.ElectionCandidate{}).Create(v).Error
}
func (i impl) UpdateById(c *gin.Context, v *entities.ElectionCandidate) error {
	return i.db.Gdb().WithContext(c).Model(&entities.ElectionCandidate{}).Where("id = ?", v.ID).Updates(v).Error
}
