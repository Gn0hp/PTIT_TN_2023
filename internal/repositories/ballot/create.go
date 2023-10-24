package ballot

import (
	"PTIT_TN/internal/entities"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(c *gin.Context, v *entities.Ballot) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Ballot{}).Create(v).Error
}
func (i impl) UpdateById(c *gin.Context, v *entities.Ballot) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Ballot{}).Where("id = ?", v.ID).Updates(v).Error
}
