package mysql

import (
	"errors"
	"fmt"
)

type Config struct {
	Host string
	Port int
	User string
	Pass string
	Name string

	Params map[string]string
}

// DSN returns a MySQL driver compatible data source name.
func (c Config) DSN() string {
	var params string
	if len(c.Params) > 0 {
		var query string
		for key, val := range c.Params {
			if query != "" {
				query += "&"
			}
			query += key + "=" + val
		}
		params = "?" + query
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Name,
		params)
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

	if c.Name == "" {
		return errors.New("database name is required")
	}

	return nil
}
