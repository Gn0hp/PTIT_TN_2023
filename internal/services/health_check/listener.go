package health_check

import (
	gosundheit "github.com/AppsFlyer/go-sundheit"
	"logur.dev/logur"
)

type healthListener struct {
	name   string
	logger logur.Logger
}

func NewHealthCheckListener(serviceName string, logger logur.Logger) gosundheit.HealthListener {
	return &healthListener{
		name:   serviceName,
		logger: logger,
	}
}

func (h healthListener) OnResultsUpdated(results map[string]gosundheit.Result) {
	for _, result := range results {
		if result.Error != nil {
			h.logger.Error("health check failed", map[string]interface{}{
				"service": h.name,
				"error":   result.Error.Error(),
				"details": result,
			})
			return
		}
	}
	h.logger.Info("health check completed", map[string]interface{}{
		"service": h.name,
	})
}
