package user

import (
	"PTIT_TN/internal/repositories"
	"logur.dev/logur"
)

type Handler struct {
	logger logur.LoggerFacade
	repo   repositories.Registry
}

func New(logger logur.LoggerFacade, registry repositories.Registry) Handler {
	return Handler{
		logger: logger,
		repo:   registry,
	}
}
