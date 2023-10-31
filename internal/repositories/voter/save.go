package voter

import (
	"PTIT_TN/internal/entities"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(ctx *gin.Context, v *entities.Voter) error {
	return i.db.Gdb().WithContext(ctx).Model(&entities.Voter{}).Create(v).Error
}

func (i impl) UpdateById(ctx *gin.Context, v *entities.Voter) error {
	return i.db.Gdb().WithContext(ctx).Model(&entities.Voter{}).Where("id = ?", v.ID).Updates(v).Error
}

func (i impl) SelfUpdate(ctx *gin.Context, v *entities.Voter) error {
	return i.db.Gdb().WithContext(ctx).Model(&entities.Voter{}).Where("id = ?", v.ID).Updates(v).Error
}

func (i impl) Delete(c *gin.Context, id uint64) error {
	return i.db.Gdb().WithContext(c).Model(&entities.Voter{}).Where("id = ?", id).Delete(&entities.Voter{}).Error
}
