package router

import (
	"github.com/gin-gonic/gin"
)

func (r *routerImpl) addEventEndpoints(apiV1 *gin.RouterGroup) {

	// GET /v1/events
	// GET /v1/events/details/:eventId
	// GET /v1/events/by-user/:userId
	// POST /v1/events/create
	// PATCH /v1/events/update
	// DELETE /v1/events/:eventId

	eventRoutes := apiV1.Group("/events")
	{
		eventRoutes.GET("", r.eventHandler.GetEvents)
		eventRoutes.GET("/details/:eventId", r.eventHandler.GetEventDetails)

		eventRoutes.GET("/by-user/:user-id", r.eventHandler.GetAllEventsByUserId)
		eventRoutes.POST("/create", r.eventHandler.CreateEvent)

		eventRoutes.PATCH("/update", r.eventHandler.UpdateEvent)
		eventRoutes.DELETE("/:eventId", r.eventHandler.UpdateEvent)
	}
}
