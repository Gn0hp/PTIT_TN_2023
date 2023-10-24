package voter

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

type Repo interface {
	Create(ctx *gin.Context, v *entities.Voter) error
	UpdateById(ctx *gin.Context, v *entities.Voter) error
	SelfUpdate(ctx *gin.Context, v *entities.Voter) error
	FindAll(c *gin.Context) ([]*entities.Voter, error)
	FindById(c *gin.Context, id uint64) (*entities.Voter, error)
	FindByCccd(c *gin.Context, cccd string) (*entities.Voter, error)
	GetPendingVoters(c *gin.Context) ([]*entities.Voter, error)
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
