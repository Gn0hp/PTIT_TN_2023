package user

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db"
	"logur.dev/logur"
)

type Repo interface {
	FindAll() ([]*entities.User, error)
	Create(u *entities.User) error
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
