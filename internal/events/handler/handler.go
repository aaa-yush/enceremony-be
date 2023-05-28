package handler

import (
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"github.com/gin-gonic/gin"
)

type impl struct {
	logger *logger.Logger
	conf   config.Config
}

type EventsHandler interface {
	GetEvents(c *gin.Context)
	GetEventDetails(c *gin.Context)
}

func NewEventsHandler(
	log *logger.Logger,
	conf config.Config,
) EventsHandler {
	return &impl{
		logger: log,
		conf:   conf,
	}
}
