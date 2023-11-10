package voter

import (
	"PTIT_TN/internal/repositories"
	"PTIT_TN/pkg/rabbitMQ"
	"logur.dev/logur"
)

type Handler struct {
	logger   logur.LoggerFacade
	repo     repositories.Registry
	rabbitMQ rabbitMQ.MqService
}

func New(logger logur.LoggerFacade, registry repositories.Registry, mq rabbitMQ.MqService) Handler {
	return Handler{
		logger:   logger,
		repo:     registry,
		rabbitMQ: mq,
	}
}
