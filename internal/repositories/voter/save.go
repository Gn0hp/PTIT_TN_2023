package voter

import (
	"PTIT_TN/internal/entities"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(ctx *gin.Context, v *entities.Voter) error {
	//TODO implement me
	return i.db.Gdb().WithContext(ctx).Model(&entities.Voter{}).Create(v).Error
}

func (i impl) UpdateById(ctx *gin.Context, v *entities.Voter) error {
	//TODO implement me
	return i.db.Gdb().WithContext(ctx).Model(&entities.Voter{}).Where("id = ?", v.ID).Updates(v).Error
}

func (i impl) SelfUpdate(ctx *gin.Context, v *entities.Voter) error {
	//TODO implement me
	return i.db.Gdb().WithContext(ctx).Model(&entities.Voter{}).Where("id = ?", v.ID).Updates(v).Error
}
