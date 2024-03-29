package ballotRoleElect

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

type Repo interface {
	Create(c *gin.Context, v *entities.BallotRoleElect) error
	CreateInBatches(c *gin.Context, v []*entities.BallotRoleElect) error
	FindAll(c *gin.Context) ([]*entities.BallotRoleElect, error)
	FindById(c *gin.Context, id uint64) (*entities.BallotRoleElect, error)
	FindByOption(c *gin.Context, option entities.BallotRoleElect) ([]*entities.BallotRoleElect, error)
	FindByBallotId(c *gin.Context, ballotId uint64) ([]*StatByBallotResponse, error)
	FindByRoleElectId(c *gin.Context, roleElectId uint64) ([]*entities.BallotRoleElect, error)
	CountVoterByCandidateId(c *gin.Context, roleElectId uint64) (int64, error)
	StatVoterByRoleElectId(c *gin.Context, roleElectId uint64) ([]*StatBreByCandidateDbFind, error)
}
type impl struct {
	logger logur.LoggerFacade
	db     *db.GormDB
}

func (i impl) Create(c *gin.Context, v *entities.BallotRoleElect) error {
	return i.db.Gdb().WithContext(c).Model(&entities.BallotRoleElect{}).Create(v).Error
}
func (i impl) CreateInBatches(c *gin.Context, v []*entities.BallotRoleElect) error {
	return i.db.Gdb().WithContext(c).Model(&entities.BallotRoleElect{}).CreateInBatches(v, 20).Error
}

func New(logger logur.LoggerFacade, db *db.GormDB) Repo {
	return &impl{
		logger: logger,
		db:     db,
	}
}
