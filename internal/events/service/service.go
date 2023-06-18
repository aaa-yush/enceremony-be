package service

import (
	"context"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"enceremony-be/internal/events/repo"
	"enceremony-be/pkg/events"
	"go.uber.org/zap"
)

type EventService interface {
	GetAllEvents(ctx context.Context) (*events.EventList, error)

	GetEventDetails(ctx context.Context, id string) (*events.EventDetails, error)
}

type impl struct {
	logger *logger.Logger
	conf   *config.Config
	repo   repo.EventsRepo
}

func (i *impl) GetAllEvents(ctx context.Context) (*events.EventList, error) {

	repoRes, err := i.repo.GetEvents(ctx)
	if err != nil {
		i.logger.Errorw("GetAllEvents", zap.String("ctx", "repoFail"), zap.Error(err))
		return nil, err
	}

	el := events.EventList{}

	for _, v := range repoRes {
		el.Events = append(el.Events, events.EventListItem{
			EventDetailCommons: events.EventDetailCommons{
				Id: v.Id,
				Creator: &events.Creator{
					Id:   v.Id,
					Name: v.Name,
				},
				CAt:      v.CreatedAt,
				UAt:      v.EventDate,
				Name:     v.Name,
				ShareUrl: "TBD",
			},
		})
	}

	return &el, err

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
		logger: log,
		conf:   conf,
		repo:   eventRepo,
	}
}
