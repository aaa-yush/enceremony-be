package router

import (
	"context"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/common/middleware"
	"enceremony-be/internal/config"
	"enceremony-be/internal/events/handler"
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
	Start(context.Context) error
}

type routerImpl struct {
	conf         *config.Config
	engine       *gin.Engine
	eventHandler handler.EventsHandler
	logger       *logger.Logger
}

func NewRouter(
	conf *config.Config,
	eventHandler handler.EventsHandler,
	logger *logger.Logger,
) Router {

	router := gin.New()

	return &routerImpl{
		conf:         conf,
		eventHandler: eventHandler,
		engine:       router,
		logger:       logger,
	}
}

func (r *routerImpl) Start(ctx context.Context) error {
	// Start async worker for publishing events.
	//r.asyncDispatcher.Run(ctx)

	fmt.Println("Swagger => http://localhost/internal/swagger/index.html")
	return r.engine.Run("0.0.0.0:8080")
}

func (r *routerImpl) MapRoutes() {

	r.engine.Use(location.Default())

	r.engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "Ence Server is running")
	})

	r.engine.Use(middleware.RecoveryWithZap(r.logger.Desugar(), true))

}
