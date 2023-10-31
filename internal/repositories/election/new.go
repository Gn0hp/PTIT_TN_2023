package election

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

type Repo interface {
	Create(c *gin.Context, e *entities.Election) error
	UpdateById(c *gin.Context, e *entities.Election) error
	FindAll(c *gin.Context) ([]*entities.Election, error)
	FindById(c *gin.Context, id uint64) (*entities.Election, error)
	FindByOption(c *gin.Context, option entities.Election) ([]*entities.Election, error)
	FindOpenElection(c *gin.Context) (*entities.Election, error)
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
