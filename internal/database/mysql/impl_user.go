package mysql

import (
	"context"
	"enceremony-be/internal/database/mysql/models"
	"fmt"
	"math/rand"
	"time"
)

func (m *mysqlStoreImpl) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	user.Id = generateUserId()

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.User{}).
		Create(&user).
		Error

	return user, err
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

func (m *mysqlStoreImpl) GetUserDetailsByEmail(ctx context.Context, email string) (*models.User, error) {

	var res models.User

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.User{}).
		Where(&models.User{Email: email}).
		First(&res).
		Error

	return &res, err
}

func generateUserId() string {

	min := 10
	max := 100
	buffer := min + rand.Intn(max-min+1)
	return fmt.Sprintf("%d%d", time.Now().Unix(), buffer)
}
