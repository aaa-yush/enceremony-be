package router

import "github.com/gin-gonic/gin"

func (r *routerImpl) addEventEndpoints(apiV1 *gin.RouterGroup) {

	eventRoutes := apiV1.Group("/events")
	{
		eventRoutes.GET("", r.eventHandler.GetEvents)
	}
}
