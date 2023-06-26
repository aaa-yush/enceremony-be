package mysql

import (
	"context"
	"enceremony-be/internal/database/mysql/models"
)

// TODO: Implement Pagination
func (m *mysqlStoreImpl) GetAllEvents(ctx context.Context) ([]models.Event, error) {
	res := []models.Event{}

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.Event{}).
		Limit(10).
		Find(&res).
		Error

	return res, err
}

func (m *mysqlStoreImpl) GetAllEventsByUserId(ctx context.Context, userId string) ([]models.Event, error) {

	res := []models.Event{}

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.Event{}).
		Where("user_id = ?", userId).
		Limit(10).
		Find(&res).
		Error

	return res, err
}

func (m *mysqlStoreImpl) GetEventDetails(ctx context.Context, eventId string) (*models.Event, error) {

	res := models.Event{}

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.Event{}).
		Where(&models.Event{Id: eventId}).
		First(&res).
		Error

	return &res, err
}

func (m *mysqlStoreImpl) UpdateEvent(ctx context.Context, updateEvent *models.Event) (*models.Event, error) {

	res := models.Event{}

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.Event{}).
		Where(&models.Event{Id: updateEvent.Id}).
		Updates(updateEvent).
		Error

	return &res, err
}

func (m *mysqlStoreImpl) InsertEvent(ctx context.Context, insertData *models.Event) error {

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.Event{}).
		Save(&insertData).
		Error

	return err
}

func (m *mysqlStoreImpl) DeleteEvent(ctx context.Context, id string) error {

	deletePredicate := models.Event{Id: id}

	return m.mysqlDb.WithContext(ctx).Model(&models.Event{}).Where(deletePredicate).Delete(&deletePredicate).Error
}
