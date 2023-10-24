package electionResult

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

type Repo interface {
	Create(c *gin.Context, v *entities.ElectionResult) error
	UpdateById(c *gin.Context, v *entities.ElectionResult) error
	FindAll(c *gin.Context) ([]*entities.ElectionResult, error)
	FindById(c *gin.Context, id uint64) (*entities.ElectionResult, error)
	FindByElectionCandidateId(c *gin.Context, electionCandidateId uint64) ([]*entities.ElectionResult, error)
	FindByOption(c *gin.Context, option entities.ElectionResult) ([]*entities.ElectionResult, error)
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
