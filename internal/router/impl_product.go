package router

import "github.com/gin-gonic/gin"

func (r *routerImpl) addProductEndpoints(apiV1 *gin.RouterGroup) {

	// GET /v1/product/123?eventId=1
	// POST /v1/product
	// DELETE /v1/product
	// PATCH /v1/product/status
	// PATCH /v1/product

	productRoutes := apiV1.Group("/product")
	{
		productRoutes.GET("/:id", r.productHandler.GetProductDetails)
		productRoutes.POST("", r.productHandler.AddProduct)
		productRoutes.DELETE("", r.productHandler.DeleteProduct)

		productRoutes.PATCH("", r.productHandler.UpdateProduct)
		productRoutes.PATCH("/status", r.productHandler.UpdateProductPurchaseStatus)
	}
}
