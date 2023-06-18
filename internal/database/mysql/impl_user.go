package mysql

import (
	"context"
	"enceremony-be/internal/database/mysql/models"
)

func (m *mysqlStoreImpl) CreateUser(ctx context.Context, user *models.User) error {

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.User{}).
		Save(&user).
		Error

	return err
}

func (m *mysqlStoreImpl) GetUserDetails(ctx context.Context, userId string) (*models.User, error) {

	res := models.User{}

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.User{}).
		Where(&models.User{Id: userId}).
		First(&res).
		Error

	return &res, err
}
