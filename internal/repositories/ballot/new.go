package ballot

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

type Repo interface {
	Create(c *gin.Context, v *entities.Ballot) error
	UpdateById(c *gin.Context, v *entities.Ballot) error
	FindAll(c *gin.Context) ([]*entities.Ballot, error)
	FindById(c *gin.Context, id uint64) (*entities.Ballot, error)
	FindByOption(c *gin.Context, option entities.Ballot) ([]*entities.Ballot, error)
	FindByVoterId(c *gin.Context, voterId uint64) ([]*entities.Ballot, error)
	FindByResultId(c *gin.Context, resultId uint64) ([]*entities.Ballot, error)
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
