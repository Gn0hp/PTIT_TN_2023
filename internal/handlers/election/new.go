package election

import (
	"PTIT_TN/internal/repositories"
	"PTIT_TN/pkg/rabbitMQ"
	"logur.dev/logur"
)

type Handler struct {
	logger    logur.LoggerFacade
	repo      repositories.Registry
	mqService rabbitMQ.MqService
}

func New(logger logur.LoggerFacade, repos repositories.Registry, service rabbitMQ.MqService) Handler {
	return Handler{
		logger:    logger,
		repo:      repos,
		mqService: service,
	}
}
