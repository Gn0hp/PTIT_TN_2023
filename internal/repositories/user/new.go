package user

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/query_params"
	"PTIT_TN/internal/services/db"
	"context"
	"logur.dev/logur"
)

type Repo interface {
	FindAll(c context.Context, params query_params.GetUserParams, lock bool) ([]*entities.User, error)
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
