package candidate

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

type Repo interface {
	Create(c *gin.Context, v *entities.Candidate) error
	CheckRegistered(c *gin.Context, cccd string, roleElectId uint64) (bool, error)
	UpdateById(c *gin.Context, id uint64, v *entities.Candidate) error
	SelfUpdate(c *gin.Context, v *entities.Candidate) error
	FindAll(c *gin.Context) ([]*entities.Candidate, error)
	FindByOption(c *gin.Context, option entities.Candidate) ([]*entities.Candidate, error)
	FindByCccd(c *gin.Context, cccd string) (*entities.Candidate, error)
	FindById(c *gin.Context, id uint64) (*entities.Candidate, error)
	FindCandidateByElectionRoleId(c *gin.Context, electionRoleId uint64) ([]*ViewCandidateDbFind, error)
	GetPendingCandidate(c *gin.Context) ([]*entities.Candidate, error)
	Delete(c *gin.Context, id uint64) error
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
