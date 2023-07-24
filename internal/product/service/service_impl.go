package service

import (
	"context"
	"enceremony-be/internal/database/mysql/models"
	"enceremony-be/pkg/product"
	"go.uber.org/zap"
)

func (i *impl) GetProductDetails(ctx context.Context, productId, eventId string) (*product.ProductDetailsResponse, error) {

	repoRes, err := i.repo.GetProductDetailsByEventId(ctx, eventId, productId)
	if err != nil {
		i.logger.Errorw("GetProductDetails", zap.String("ctx", "GetProductDetailsByEventId"),
			zap.Error(err))
		return nil, err
	}

	res := product.ProductDetailsResponse{
		Id:          repoRes.Id,
		Name:        repoRes.Name,
		UserId:      repoRes.UserId,
		ImageUrl:    "TBD",
		Description: repoRes.Description,
		Link:        repoRes.Link,
		EventId:     eventId,
		IsPurchased: repoRes.IsPurchased,
	}

	return &res, nil
}

func (i *impl) AddProduct(ctx context.Context, productDetail *product.CreateProductRequest) error {

	createReq := models.Product{
		Name:        productDetail.Name,
		UserId:      productDetail.UserId,
		Description: productDetail.Description,
		Link:        productDetail.Link,
	}

	err := i.repo.AddProduct(ctx, &createReq)
	if err != nil {
		i.logger.Errorw("AddProduct", zap.String("ctx", "Repo"),
			zap.Error(err), zap.Any("request", productDetail))
		return err
	}

	return nil
}

func (i *impl) DeleteProduct(ctx context.Context, req *product.DeleteRequest) error {

	return i.repo.DeleteProduct(ctx, req.Id)
}

func (i *impl) UpdateProductPurchaseStatus(ctx context.Context, req *product.UpdateProductPurchaseStatus) error {
	return i.repo.UpdateProductPurchaseStatus(ctx, req.ProductId, req.EventId, req.IsPurchased)
}

func (i *impl) EditProduct(ctx context.Context, req *product.EditRequest) (*product.ProductDetailsResponse, error) {

	udpateReq := models.Product{
		Id:          req.Id,
		Name:        req.Name,
		UserId:      req.UserId,
		Description: req.Description,
		Link:        req.Link,
	}

	err := i.repo.EditProduct(ctx, &udpateReq)
	if err != nil {
		i.logger.Errorw("AddProduct", zap.String("ctx", "Repo"),
			zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return i.GetProductDetails(ctx, req.Id, req.EventId)
}
