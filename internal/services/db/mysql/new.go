package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnector(c Config) (*gorm.DB, error) {
	c.Params["parseTime"] = "true"
	c.Params["rejectReadOnly"] = "true"

	connectString := c.DSN()
	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
