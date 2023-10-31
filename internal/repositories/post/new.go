package post

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

type Repo interface {
	Create(c *gin.Context, v *entities.Post) error
	UpdateById(c *gin.Context, v *entities.Post) error
	FindAll(c *gin.Context) ([]*entities.Post, error)
	FindById(c *gin.Context, id uint64) (*entities.Post, error)
	FindByCandidateId(c *gin.Context, candidateId uint64) ([]*entities.Post, error)
	FindByOption(c *gin.Context, option entities.Post) ([]*entities.Post, error)
	Delete(c *gin.Context, id uint64) error
	Update(c *gin.Context, id uint64, v *entities.Post) error
}
type impl struct {
	logger logur.LoggerFacade
	db     *db.GormDB
}

func New(logger logur.LoggerFacade, db *db.GormDB) Repo {
	return &impl{
		logger: logger,
		db:     db,
	}
}
