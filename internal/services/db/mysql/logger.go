package mysql

import (
	"github.com/go-sql-driver/mysql"
	"logur.dev/logur"
)

func SetLogger(logger logur.Logger) {
	logger = logur.WithField(logger, "component", "mysql")

	_ = mysql.SetLogger(logur.NewErrorPrintLogger(logger))
}
