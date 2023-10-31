package electionRole

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

type Repo interface {
	Create(c *gin.Context, v *entities.ElectionRole) error
	BulkInsert(c *gin.Context, v []*entities.ElectionRole) error
	CreateMulti(c *gin.Context, v []*entities.ElectionRole) error
	UpdateById(c *gin.Context, v *entities.ElectionRole) error
	FindAll(c *gin.Context) ([]*entities.ElectionRole, error)
	FindById(c *gin.Context, id uint64) (*entities.ElectionRole, error)
	FindByOption(c *gin.Context, option entities.ElectionRole) ([]*entities.ElectionRole, error)
	FindByElectionCandidateId(c *gin.Context, electionCandidateId uint64) ([]*entities.ElectionRole, error)
	FindByElectionId(c *gin.Context, electionId uint64) ([]*entities.ElectionRole, error)
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
