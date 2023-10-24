package admin

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"logur.dev/logur"
)

type Repo interface {
	FindByUsername(c *gin.Context, username string) (*entities.Admin, error)
}

type impl struct {
	logger logur.LoggerFacade
	db     *db.GormDB
}

func (i impl) FindByUsername(c *gin.Context, username string) (*entities.Admin, error) {
	admin := &entities.Admin{}
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Admin{}).
		Where("username = ?", username).Where("status = ?", "ACTIVE")
	err := query.First(admin).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find Candidate failed, detail: %v", err))
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("admin not found")
		}
		return nil, err
	}
	return admin, nil
}

func New(logger logur.LoggerFacade, db *db.GormDB) Repo {
	return &impl{
		logger: logger,
		db:     db,
	}
}
