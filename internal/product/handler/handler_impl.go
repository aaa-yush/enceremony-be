package handler

import (
	"enceremony-be/pkg/commons"
	"enceremony-be/pkg/product"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *impl) GetProductDetails(c *gin.Context) {

	eventId := c.Query("eventId")
	if eventId == "" {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  "EventIdMissing",
		})
		return
	}

	productId := c.Param("id")
	if productId == "" {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  "ProductIdMissing",
		})
		return
	}

	resp, err := i.prodSvc.GetProductDetails(c, productId, eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (i *impl) AddProduct(c *gin.Context) {
	createReq := product.CreateProductRequest{}
	err := c.Bind(&createReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	err = i.prodSvc.AddProduct(c, &createReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commons.GenericSuccessResponse{Status: "success"})
}

func (i *impl) DeleteProduct(c *gin.Context) {
	delReq := product.DeleteRequest{}
	err := c.Bind(&delReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	err = i.prodSvc.DeleteProduct(c, &delReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commons.GenericSuccessResponse{Status: "success"})
}

func (i *impl) UpdateProductPurchaseStatus(c *gin.Context) {
	editReq := product.UpdateProductPurchaseStatus{}
	err := c.Bind(&editReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	err = i.prodSvc.UpdateProductPurchaseStatus(c, &editReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commons.GenericSuccessResponse{Status: "success"})
}

func (i *impl) UpdateProduct(c *gin.Context) {

	editReq := product.EditRequest{}
	err := c.Bind(&editReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	res, err := i.prodSvc.EditProduct(c, &editReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
