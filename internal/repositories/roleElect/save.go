package roleElect

import (
	"PTIT_TN/internal/entities"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(c *gin.Context, v *entities.RoleElect) error {
	return i.db.Gdb().WithContext(c).Model(&entities.RoleElect{}).Create(v).Error
}

func (i impl) UpdateById(c *gin.Context, v *entities.RoleElect) error {
	return i.db.Gdb().WithContext(c).Model(&entities.RoleElect{}).Where("id = ?", v.ID).Updates(v).Error
}
