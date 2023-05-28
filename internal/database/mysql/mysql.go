package mysql

import (
	"context"
	"enceremony-be/commons/clients/mysql"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"enceremony-be/internal/database/mysql/models"
	"gorm.io/gorm"
)

type MysqlStore interface {

	// Events

	// GetAllEvents TODO: Implement Pagination
	GetAllEvents(ctx context.Context) ([]models.Event, error)
}

type mysqlStoreImpl struct {
	logger     *logger.Logger
	conf       *config.Config
	mysqlDb    *gorm.DB
	connection mysql.Connection
}

func (m *mysqlStoreImpl) Save(ctx context.Context, data interface{}) error {
	return m.mysqlDb.WithContext(ctx).Save(data).Error
}

func NewMysqlStore(logger *logger.Logger, connection mysql.Connection,
	conf *config.Config) MysqlStore {

	return &mysqlStoreImpl{
		logger:  logger,
		conf:    conf,
		mysqlDb: connection.GetInstance(),
	}
}
