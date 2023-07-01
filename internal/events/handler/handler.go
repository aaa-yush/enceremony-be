package handler

import (
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"enceremony-be/internal/events/service"
	"github.com/gin-gonic/gin"
)

type impl struct {
	logger   *logger.Logger
	conf     *config.Config
	eventSvc service.EventService
}

type EventsHandler interface {
	GetEvents(c *gin.Context)
	GetEventDetails(c *gin.Context)

	CreateEvent(c *gin.Context)

	GetAllEventsByUserId(c *gin.Context)

	UpdateEvent(c *gin.Context)
}

func NewEventsHandler(
	log *logger.Logger,
	conf *config.Config,
	eventSvc service.EventService,
) EventsHandler {
	return &impl{
		logger:   log,
		conf:     conf,
		eventSvc: eventSvc,
	}
}
