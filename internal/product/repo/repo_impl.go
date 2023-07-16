package repo

import (
	"context"
	"enceremony-be/internal/database/mysql/models"
)

func (i *impl) GetProductsByEventId(ctx context.Context, eventId string) ([]models.EventProductDetails, error) {
	return i.mysqlStore.GetProductsByEventId(ctx, eventId)
}

func (i *impl) GetProductDetails(ctx context.Context, productId string) (*models.Product, error) {
	return i.mysqlStore.GetProductDetails(ctx, productId)
}

func (i *impl) AddProduct(ctx context.Context, productDetail *models.Product) error {
	return i.mysqlStore.AddProduct(ctx, productDetail)
}

func (i *impl) DeleteProduct(ctx context.Context, id string) error {
	return i.mysqlStore.DeleteProduct(ctx, id)
}

func (i *impl) UpdateProductPurchaseStatus(ctx context.Context, productId, eventId string, status bool) error {
	return i.mysqlStore.UpdateProductPurchasedStatus(ctx, productId, eventId, status)
}

func (i *impl) GetProductDetailsByEventId(ctx context.Context, eventId, productId string) (*models.EventProductDetails, error) {
	return i.mysqlStore.GetProductDetailsByEventId(ctx, eventId, productId)
}

func (i *impl) EditProduct(ctx context.Context, productDetail *models.Product) error {

	return i.mysqlStore.UpdateProduct(ctx, productDetail)
}
