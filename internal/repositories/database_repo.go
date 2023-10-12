package repositories

import (
	"PTIT_TN/internal/repositories/user"
	"PTIT_TN/internal/services/db"
	"logur.dev/logur"
)

type DatabaseRepo interface {
	User() user.Repo
}
type dbImpl struct {
	user user.Repo
}

func (d dbImpl) User() user.Repo {
	return d.user
}

func NewDatabaseRepo(logger logur.LoggerFacade, db *db.GormDB) DatabaseRepo {
	return &dbImpl{
		user: user.New(logger, db),
	}
}
