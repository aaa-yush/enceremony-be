package repo

import (
	"context"
	"enceremony-be/internal/config"
	"enceremony-be/internal/database/mysql"
	"enceremony-be/internal/database/mysql/models"
)

type impl struct {
	mysqlStore mysql.MysqlStore
	conf       *config.Config
}

type ProductRepo interface {
	GetProductsByEventId(ctx context.Context, eventId string) ([]models.EventProductDetails, error)
	GetProductDetailsByEventId(ctx context.Context, eventId, productId string) (*models.EventProductDetails, error)

	GetProductDetails(ctx context.Context, productId string) (*models.Product, error)
	AddProduct(ctx context.Context, productDetail *models.Product) error
	DeleteProduct(ctx context.Context, id string) error

	EditProduct(ctx context.Context, productDetail *models.Product) error

	UpdateProductPurchaseStatus(ctx context.Context, productId, eventId string, status bool) error
}

func NewProductRepo(
	conf *config.Config,
	mysqlStore mysql.MysqlStore) ProductRepo {
	return &impl{
		conf:       conf,
		mysqlStore: mysqlStore,
	}
}
