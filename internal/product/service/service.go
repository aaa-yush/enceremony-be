package service

import (
	"context"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	repo2 "enceremony-be/internal/product/repo"
	"enceremony-be/pkg/product"
)

type impl struct {
	logger *logger.Logger
	conf   *config.Config
	repo   repo2.ProductRepo
}

type ProductService interface {
	GetProductDetails(ctx context.Context, productId, eventId string) (*product.ProductDetailsResponse, error)

	AddProduct(ctx context.Context, productDetail *product.CreateProductRequest) error

	DeleteProduct(ctx context.Context, req *product.DeleteRequest) error

	EditProduct(ctx context.Context, req *product.EditRequest) (*product.ProductDetailsResponse, error)

	UpdateProductPurchaseStatus(ctx context.Context, req *product.UpdateProductPurchaseStatus) error
}

func NewProductService(
	logger *logger.Logger,
	conf *config.Config,
	repo repo2.ProductRepo) ProductService {
	return &impl{
		logger: logger,
		conf:   conf,
		repo:   repo,
	}
}
