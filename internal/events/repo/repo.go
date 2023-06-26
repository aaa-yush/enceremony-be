package repo

import (
	"context"
	"enceremony-be/internal/config"
	"enceremony-be/internal/database/mysql"
	"enceremony-be/internal/database/mysql/models"
)

type EventsRepo interface {
	GetEvents(ctx context.Context) ([]models.Event, error)

	GetEventDetails(ctx context.Context, eventId string) (*models.EventDetails, error)

	InsertEvent(ctx context.Context, insertData *models.Event) error
	GetAllEventsByUserId(ctx context.Context, userId string) ([]models.Event, error)
	UpdateEvent(ctx context.Context, updateEvent *models.Event) (*models.Event, error)
}

type eventsRepoImpl struct {
	conf       *config.Config
	mysqlStore mysql.MysqlStore
}

func (e *eventsRepoImpl) InsertEvent(ctx context.Context, insertData *models.Event) error {
	return e.mysqlStore.InsertEvent(ctx, insertData)
}

func (e *eventsRepoImpl) GetAllEventsByUserId(ctx context.Context, userId string) ([]models.Event, error) {
	return e.mysqlStore.GetAllEventsByUserId(ctx, userId)
}

func (e *eventsRepoImpl) UpdateEvent(ctx context.Context, updateEvent *models.Event) (*models.Event, error) {
	return e.mysqlStore.UpdateEvent(ctx, updateEvent)
}

func NewEventsRepo(
	conf *config.Config,
	mysqlStore mysql.MysqlStore) EventsRepo {
	return &eventsRepoImpl{
		conf:       conf,
		mysqlStore: mysqlStore,
	}
}

func (e *eventsRepoImpl) GetEvents(ctx context.Context) ([]models.Event, error) {
	return e.mysqlStore.GetAllEvents(ctx)
}

func (e *eventsRepoImpl) GetEventDetails(ctx context.Context, eventId string) (*models.EventDetails, error) {

	products, err := e.mysqlStore.GetProductsByEventId(ctx, eventId)
	if err != nil {
		return nil, err
	}

	ed, err := e.mysqlStore.GetEventDetails(ctx, eventId)

	resp := models.EventDetails{
		Event:    *ed,
		Products: products,
	}

	return &resp, nil
}
