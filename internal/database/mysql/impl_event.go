package mysql

import (
	"context"
	"enceremony-be/internal/database/mysql/models"
)

func (m *mysqlStoreImpl) GetAllEvents(ctx context.Context) ([]models.Event, error) {
	res := []models.Event{}
	err := m.mysqlDb.WithContext(ctx).Model(&models.Event{}).Limit(10).Find(&res).Error
	return res, err
}
