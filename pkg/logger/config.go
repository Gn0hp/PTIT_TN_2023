package log

// Config represents the configuration of a logger. Modify in `config.toml` file.
type Config struct {
	// Format specifies the output format of the logger. accepted values: json, logfmt
	Format string `mapstructure:"format"`
	// Level is the minimum log level that should appear on the output. accepted values: trace, debug, info, warn, error, fatal, panic
	Level string `mapstructure:"level"`
	// NoColor makes sure that no log output gets colorized.
	NoColor bool `mapstructure:"no_color"`
	// FullTimestamp makes sure that the full timestamp should appear on output.
	FullTimestamp bool `mapstructure:"full_timestamp"`
}
