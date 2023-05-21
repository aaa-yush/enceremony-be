package service

import (
	"context"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"enceremony-be/internal/events/repo"
	"enceremony-be/pkg/events"
)

type EventService interface {
	GetAllEvents(ctx context.Context) (*events.EventList, error)

	GetEventDetails(ctx context.Context, id string) (*events.EventDetails, error)
}

type impl struct {
	log  *logger.Logger
	conf *config.Config
	repo repo.EventsRepo
}

func (i *impl) GetAllEvents(ctx context.Context) (*events.EventList, error) {
	//TODO implement me
	panic("implement me")
}

func (i *impl) GetEventDetails(ctx context.Context, id string) (*events.EventDetails, error) {
	//TODO implement me
	panic("implement me")
}

func NewEventService(
	log *logger.Logger,
	conf *config.Config,
	eventRepo repo.EventsRepo,
) EventService {
	return &impl{
		log:  log,
		conf: conf,
		repo: eventRepo,
	}
}
