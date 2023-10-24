package roleElect

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

type Repo interface {
	Create(c *gin.Context, v *entities.RoleElect) error
	UpdateById(c *gin.Context, v *entities.RoleElect) error
	FindAll(c *gin.Context) ([]*entities.RoleElect, error)
	FindById(c *gin.Context, id uint64) (*entities.RoleElect, error)
	FindByCandidateId(c *gin.Context, candidateId uint64) ([]*entities.RoleElect, error)
	FindByElectionCandidateId(c *gin.Context, electionCandidateId uint64) ([]*entities.RoleElect, error)
	FindByOption(c *gin.Context, option entities.RoleElect) ([]*entities.RoleElect, error)
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