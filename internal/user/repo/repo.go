package repo

import (
	"context"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"enceremony-be/internal/database/mysql"
	"enceremony-be/internal/database/mysql/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetOrCreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type repoImpl struct {
	conf       *config.Config
	logger     *logger.Logger
	mysqlStore mysql.MysqlStore
}

func NewUserRepo(
	conf *config.Config,
	logger *logger.Logger,
	mysqlStore mysql.MysqlStore,
) UserRepo {
	return &repoImpl{
		conf:       conf,
		logger:     logger,
		mysqlStore: mysqlStore,
	}
}

func (r *repoImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return r.mysqlStore.GetUserDetailsByEmail(ctx, email)
}

func (r *repoImpl) GetOrCreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	userDetails, err := r.mysqlStore.GetUserDetailsByEmail(ctx, user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		userDetails, err := r.mysqlStore.CreateUser(ctx, user)
		if err != nil {
			return nil, err
		}

		return userDetails, nil
	}

	return userDetails, err
}
