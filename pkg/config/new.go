package config

import (
	"PTIT_TN/internal/services/db/mysql"
	log "PTIT_TN/pkg/logger"
	"PTIT_TN/pkg/rabbitMQ"
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

type appConfig struct {
	HttpAddr string `mapstructure:"httpAddr"` //HttpServer address
	Storage  string // Storage is the storage backend of the application [inmemory, database]
}
type Configuration struct {
	App      appConfig  // App configuration
	Log      log.Config // Log logger config
	Database mysql.Config
	RabbitMq rabbitMQ.Config
}

func (c appConfig) Validate() error {
	if c.HttpAddr == "" {
		return errors.New("http app server address is required")
	}

	if c.Storage != "inmemory" && c.Storage != "database" {
		return errors.New("app storage must be inmemory or database")
	}

	return nil
}
func (c Configuration) Validate() error {

	if err := c.App.Validate(); err != nil {
		return err
	}

	if err := c.Database.Validate(); err != nil {
		return err
	}
	if err := c.RabbitMq.Validate(); err != nil {
		return err
	}

	return nil
}

func New(v *viper.Viper, f *pflag.FlagSet) Configuration {
	bindDefault(v, f)
	f.String("config", "", "Configuration file")
	_ = f.Parse(os.Args[1:])

	if c, _ := f.GetString("config"); c != "" {
		v.SetConfigFile(c)
	}

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to read configuration: %v", err))
	}
	notFoundErr, configFileNotFound := err.(viper.ConfigFileNotFoundError)
	if configFileNotFound {
		panic(fmt.Sprintf("config file not found: %v", notFoundErr.Error()))
	}
	var conf Configuration
	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}
	return conf
}
