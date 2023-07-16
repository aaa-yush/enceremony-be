package handler

import (
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"enceremony-be/internal/product/service"
	"github.com/gin-gonic/gin"
)

type impl struct {
	logger  *logger.Logger
	conf    *config.Config
	prodSvc service.ProductService
}

type ProductHandler interface {
	GetProductDetails(c *gin.Context)

	AddProduct(c *gin.Context)

	DeleteProduct(c *gin.Context)

	UpdateProduct(c *gin.Context)

	UpdateProductPurchaseStatus(c *gin.Context)
}

func NewProductHandler(
	logger *logger.Logger,
	conf *config.Config,
	prodSvc service.ProductService) ProductHandler {
	return &impl{
		logger:  logger,
		conf:    conf,
		prodSvc: prodSvc,
	}
}
