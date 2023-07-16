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

	InsertEvent(ctx context.Context, insertData *models.Event) error
	GetAllEventsByUserId(ctx context.Context, userId string) ([]models.Event, error)
	GetEventDetails(ctx context.Context, eventId string) (*models.Event, error)
	UpdateEvent(ctx context.Context, updateEvent *models.Event) (*models.Event, error)
	DeleteEvent(ctx context.Context, id string) error

	/*
		USER
	*/
	CreateUser(ctx context.Context, user *models.User) error
	GetUserDetails(ctx context.Context, userId string) (*models.User, error)

	/*
		PRODUCT
	*/
	GetProductDetails(ctx context.Context, productId string) (*models.Product, error)
	AddProduct(ctx context.Context, productDetail *models.Product) error
	UpdateProduct(ctx context.Context, productDetail *models.Product) error
	DeleteProduct(ctx context.Context, id string) error

	/*
		EVENT-Product
	*/
	GetProductsByEventId(ctx context.Context, eventId string) ([]models.EventProductDetails, error)
	UpdateProductPurchasedStatus(ctx context.Context, eventId, productId string, status bool) error
	GetProductDetailsByEventId(ctx context.Context, eventId, productId string) (*models.EventProductDetails, error)
}

type mysqlStoreImpl struct {
	logger  *logger.Logger
	conf    *config.Config
	mysqlDb *gorm.DB
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
