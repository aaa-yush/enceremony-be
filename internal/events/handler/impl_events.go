package handler

import (
	"enceremony-be/pkg/commons"
	"enceremony-be/pkg/events"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *impl) GetEvents(c *gin.Context) {

	i.logger.Info("GetEvents")

	res, err := i.eventSvc.GetAllEvents(c)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (i *impl) GetEventDetails(c *gin.Context) {

	eventId := c.Param("eventId")
	if eventId == "" {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  "EventIdMissing",
		})
		return
	}

	res, err := i.eventSvc.GetEventDetails(c, eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)

}

func (i *impl) CreateEvent(c *gin.Context) {

	req := events.EventDetails{}

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	err = i.eventSvc.InsertEvent(c, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commons.GenericSuccessResponse{Status: "success"})
}

func (i *impl) GetAllEventsByUserId(c *gin.Context) {

	eventId := c.Param("userId")
	if eventId == "" {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  "EventIdMissing",
		})
		return
	}

	res, err := i.eventSvc.GetAllEventsByUserId(c, eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (i *impl) UpdateEvent(c *gin.Context) {

	updateEvent := events.EventDetails{}
	err := c.Bind(&updateEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	res, err := i.eventSvc.UpdateEvent(c, &updateEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.GenericErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
