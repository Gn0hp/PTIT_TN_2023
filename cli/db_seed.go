package main

import (
	"PTIT_TN/internal"
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/services/db/mysql"
	log "PTIT_TN/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"logur.dev/logur"
)

type MigrateService struct {
	internal.DefaultService
	logger logur.LoggerFacade
	gormDb *gorm.DB
}

func main() {
	var migrateService MigrateService
	migrateService.Init()

	table := []interface{}{
		entities.User{},
	}
	err := migrateService.gormDb.AutoMigrate(table...)
	if err != nil {
		migrateService.logger.Error(fmt.Sprintf("Seed failed, details: %v", err))
		return
	}
	migrateService.logger.Info("Seed completed")
}

func (s *MigrateService) Init() {
	s.DefaultService.Init()
	var (
		logCf = log.Config{}
		dbCf  = mysql.Config{}
	)
	cfBytes, _ := json.Marshal(viper.GetStringMap("log"))
	json.Unmarshal(cfBytes, &logCf)
	cfBytes, _ = json.Marshal(viper.GetStringMap("database"))
	json.Unmarshal(cfBytes, &dbCf)
	logger := log.NewLogger(logCf)

	if dbCf.Params == nil {
		dbCf.Params = make(map[string]string)
	}
	gormDb, err := mysql.NewConnector(dbCf)
	if err != nil {
		panic(err)
	}
	logger.Info("Database connected")
	s.logger = logger
	s.gormDb = gormDb
}
