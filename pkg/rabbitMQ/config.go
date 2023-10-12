package rabbitMQ

import "errors"

type Config struct {
	User     string
	Password string
	Host     string
	Port     int
}

// Validate checks that the configuration is valid.
func (c Config) Validate() error {
	if c.Host == "" {
		return errors.New("database host is required")
	}

	if c.Port == 0 {
		return errors.New("database port is required")
	}

	if c.User == "" {
		return errors.New("database user is required")
	}

	return nil
}
