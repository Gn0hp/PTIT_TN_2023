package electionCandidate

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

type Repo interface {
	Create(c *gin.Context, v *entities.ElectionCandidate) error
	UpdateById(c *gin.Context, v *entities.ElectionCandidate) error
	FindAll(c *gin.Context) ([]*entities.ElectionCandidate, error)
	FindById(c *gin.Context, id uint64) (*entities.ElectionCandidate, error)
	FindByElectionId(c *gin.Context, electionId uint64) ([]*entities.ElectionCandidate, error)
	FindByOption(c *gin.Context, option entities.ElectionCandidate) ([]*entities.ElectionCandidate, error)
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
