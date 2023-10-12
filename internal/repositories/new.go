package repositories

import (
	"PTIT_TN/internal/services/db"
	"PTIT_TN/internal/services/db/redis"
	"logur.dev/logur"
)

type Registry interface {
	Database() DatabaseRepo
	Cache() CacheRepo
}

type impl struct {
	cacheRepo CacheRepo
	dbRepo    DatabaseRepo
}

func New(logger logur.LoggerFacade, db *db.GormDB, client *redis.Client) Registry {
	return &impl{
		cacheRepo: NewCacheRepo(client),
		dbRepo:    NewDatabaseRepo(logger, db),
	}
}

func (i impl) Database() DatabaseRepo {
	return i.dbRepo
}

func (i impl) Cache() CacheRepo {
	return i.cacheRepo
}
