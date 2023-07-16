package mysql

import (
	"context"
	"enceremony-be/internal/database/mysql/models"
)

func (m *mysqlStoreImpl) AddProduct(ctx context.Context, productDetail *models.Product) error {

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.Product{}).
		Save(&productDetail).
		Error

	return err
}

func (m *mysqlStoreImpl) GetProductDetails(ctx context.Context, productId string) (*models.Product, error) {

	res := models.Product{}

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.Product{}).
		Where(&models.Product{Id: productId}).
		First(&res).
		Error

	return &res, err
}

func (m *mysqlStoreImpl) GetProductsByEventId(ctx context.Context, eventId string) ([]models.EventProductDetails, error) {

	sql := "select p.*, ep.is_purchased from product p left join event_products ep on p.id = ep.product_id where ep.event_id = ?"

	res := []models.EventProductDetails{}

	err := m.mysqlDb.WithContext(ctx).
		Raw(sql, eventId).
		Scan(&res).
		Error

	return res, err
}

func (m *mysqlStoreImpl) GetProductDetailsByEventId(ctx context.Context, eventId, productId string) (*models.EventProductDetails, error) {

	sql := "select p.*, ep.is_purchased from product p inner join event_products " +
		"ep on p.id = ep.product_id where ep.event_id = ? and p.product_id= ?"
	res := models.EventProductDetails{}

	err := m.mysqlDb.WithContext(ctx).
		Raw(sql, eventId, productId).
		First(&res).
		Error

	return &res, err
}

func (m *mysqlStoreImpl) DeleteProduct(ctx context.Context, id string) error {

	deletePredicate := models.Product{Id: id}

	return m.mysqlDb.WithContext(ctx).Model(&models.Product{}).Where(deletePredicate).Delete(&deletePredicate).Error
}

func (m *mysqlStoreImpl) UpdateProduct(ctx context.Context, productDetail *models.Product) error {

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.Product{}).
		Where(&models.Product{Id: productDetail.Id}).
		Updates(productDetail).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlStoreImpl) UpdateProductPurchasedStatus(ctx context.Context, eventId, productId string, status bool) error {

	err := m.mysqlDb.WithContext(ctx).
		Model(&models.EventProducts{}).
		Where(&models.EventProducts{EventId: eventId, ProductId: productId}).
		Update("is_purchased", status).
		Error

	return err
}
