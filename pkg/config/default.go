package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

func bindDefault(v *viper.Viper, f *pflag.FlagSet) {
	// Set Viper defaults config
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("toml")

	// Environment variable settings
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")) // Support variable calling in the code
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	// Global configuration
	v.SetDefault("shutdownTimeout", 15*time.Second) // Graceful shutdown timeout

	// Log configuration
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		v.SetDefault("no_color", false) // Disable color output
	}
	v.SetDefault("log.format", "logfmt")
	v.SetDefault("log.level", "info")
	v.RegisterAlias("log.noColor", "no_color")

	v.SetDefault("rabbit_mq.user", "guest")
	v.SetDefault("rabbit_mq.password", "guest")
	v.SetDefault("rabbit_mq.host", "localhost")
	v.SetDefault("rabbit_mq.port", 5672)

	// Server address configuration
	f.String("http-addr", "8400", "App HTTP server address")
	_ = v.BindPFlag("app.httpAddr", f.Lookup("http-addr"))
	v.SetDefault("app.httpAddr", "8400")
	v.SetDefault("app.storage", "database")

	// Database configuration - bind to environment variables
	_ = v.BindEnv("database.host")
	v.SetDefault("database.port", 3306)
	_ = v.BindEnv("database.user")
	_ = v.BindEnv("database.pass")
	_ = v.BindEnv("database.name")

	v.SetDefault("database.params", map[string]string{
		"collation": "utf8mb4_general_ci",
	})
}
