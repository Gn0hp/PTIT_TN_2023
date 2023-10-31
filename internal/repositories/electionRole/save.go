package electionRole

import (
	"PTIT_TN/internal/entities"
	"github.com/gin-gonic/gin"
)

func (i impl) Create(c *gin.Context, v *entities.ElectionRole) error {
	return i.db.Gdb().WithContext(c).Model(&entities.ElectionRole{}).Create(v).Error
}
func (i impl) BulkInsert(c *gin.Context, v []*entities.ElectionRole) error {
	return i.db.Gdb().WithContext(c).Model(&entities.ElectionRole{}).CreateInBatches(v, 20).Error
}
func (i impl) UpdateById(c *gin.Context, v *entities.ElectionRole) error {
	return i.db.Gdb().WithContext(c).Model(&entities.ElectionRole{}).Where("id = ?", v.ID).Updates(v).Error
}

func (i impl) CreateMulti(c *gin.Context, v []*entities.ElectionRole) error {
	return i.db.Gdb().WithContext(c).Model(&entities.ElectionRole{}).Save(v).Error
}
