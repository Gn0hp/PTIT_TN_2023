package internal

import "github.com/spf13/viper"

var initiated = false

func init() {
	if !initiated {
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")
		viper.ReadInConfig()
		initiated = true
	}
}

type DefaultService struct{}

func (s *DefaultService) Init() {}
