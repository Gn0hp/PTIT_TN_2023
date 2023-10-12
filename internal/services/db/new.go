package db

import (
	"PTIT_TN/internal"
	"PTIT_TN/internal/services/db/mysql"
	"PTIT_TN/internal/services/health_check"
	"PTIT_TN/pkg"
	"fmt"
	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/AppsFlyer/go-sundheit/checks"
	"gorm.io/gorm"
	"logur.dev/logur"
	"time"
)

type GormDB struct {
	gormDb *gorm.DB
}

func New(logger logur.LoggerFacade, config mysql.Config) *GormDB {
	healthcheck := health_check.NewHealthCheckListener("mysql", logger)

	mysql.SetLogger(logger)

	gormDb, err := mysql.NewConnector(config)
	if err != nil {
		panic(fmt.Sprintf("connect database failed, error: %v", err))
	}
	sqlDb, _ := gormDb.DB()

	healthchecker := gosundheit.New(gosundheit.WithHealthListeners(healthcheck))
	_ = healthchecker.RegisterCheck(
		checks.Must(checks.NewPingCheck("mysql_db.check", sqlDb)),
		gosundheit.ExecutionPeriod(time.Minute*2))
	internal.RegisterApp(pkg.MySQLConnectorAppName.ToString(), gormDb)
	logger.Info("connected to database!")
	return &GormDB{
		gormDb: gormDb,
	}
}

func (db *GormDB) Gdb() *gorm.DB {
	return db.gormDb
}

func (db *GormDB) Close() {
	dbConnection, _ := db.gormDb.DB()
	if dbConnection != nil {
		dbConnection.Close()
	}
}
