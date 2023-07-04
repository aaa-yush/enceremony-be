package service

import (
	"context"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"enceremony-be/internal/events/repo"
	"enceremony-be/pkg/events"
	"enceremony-be/pkg/product"
	"go.uber.org/zap"
)

type EventService interface {
	GetAllEvents(ctx context.Context) (*events.EventList, error)

	GetEventDetails(ctx context.Context, id string) (*events.EventDetails, error)

	InsertEvent(ctx context.Context, insertData *events.EventDetails) error

	GetAllEventsByUserId(ctx context.Context, userId string) (*events.EventList, error)

	UpdateEvent(ctx context.Context, updateEvent *events.EventDetails) (*events.EventDetails, error)

	DeleteEvent(ctx context.Context, eventId string) error
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

	res, err := i.repo.GetEventDetails(ctx, id)
	if err != nil {
		return nil, err
	}

	products := []product.Product{}

	for _, p := range res.Products {
		products = append(products, product.Product{
			Id:       p.Id,
			Name:     p.Name,
			ImageUrl: "TBD",
		})
	}

	svcResp := events.EventDetails{
		EventDetailCommons: events.EventDetailCommons{
			Id: res.Id,
			Creator: &events.Creator{
				Id:   res.UserId,
				Name: "TBD",
			},
			CAt:       res.CreatedAt,
			UAt:       res.UpdatedAt,
			EventDate: res.EventDate,
			Name:      res.Name,
			ShareUrl:  "TBD",
		},
		PhotoUrl:         "TBD",
		Products:         products,
		EventDescription: res.Description,
	}

	return &svcResp, nil
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
