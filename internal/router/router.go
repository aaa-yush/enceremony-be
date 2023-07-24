package router

import (
	"context"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/common/middleware"
	"enceremony-be/internal/config"
	"enceremony-be/internal/events/handler"
	handler2 "enceremony-be/internal/product/handler"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
	Start(context.Context) error
}

type routerImpl struct {
	conf           *config.Config
	engine         *gin.Engine
	eventHandler   handler.EventsHandler
	productHandler handler2.ProductHandler
	logger         *logger.Logger
}

func NewRouter(
	conf *config.Config,
	eventHandler handler.EventsHandler,
	logger *logger.Logger,
	productHandler handler2.ProductHandler,
) Router {

	router := gin.New()

	return &routerImpl{
		conf:           conf,
		eventHandler:   eventHandler,
		engine:         router,
		logger:         logger,
		productHandler: productHandler,
	}
}

func (r *routerImpl) Start(ctx context.Context) error {

	return r.engine.Run("0.0.0.0:8080")
}

func (r *routerImpl) MapRoutes() {

	r.engine.Use(location.Default())

	r.engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "Ence Server is running")
	})

	apiV1 := r.engine.Group("v1")
	r.addEventEndpoints(apiV1)

	r.engine.Use(middleware.RecoveryWithZap(r.logger.Desugar(), true))

}
