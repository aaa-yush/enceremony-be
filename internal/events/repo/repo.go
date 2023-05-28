package repo

import (
	"context"
	"enceremony-be/internal/config"
	"enceremony-be/internal/database/mysql"
	"enceremony-be/internal/database/mysql/models"
)

type EventsRepo interface {
	GetEvents(ctx context.Context) ([]models.Event, error)
}

type eventsRepoImpl struct {
	conf       *config.Config
	mysqlStore mysql.MysqlStore
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
