package service

import (
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"enceremony-be/internal/events/repo"
)

type impl struct {
	logger *logger.Logger
	conf   *config.Config
	repo   repo.EventsRepo
}
