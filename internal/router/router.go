package router

import (
	"context"
	handler3 "enceremony-be/internal/auth/handler"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/common/middleware"
	"enceremony-be/internal/config"
	"enceremony-be/internal/events/handler"
	handler2 "enceremony-be/internal/product/handler"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"html/template"
)

type Router interface {
	MapRoutes()
	Start(context.Context) error
}

type routerImpl struct {
	conf              *config.Config
	engine            *gin.Engine
	eventHandler      handler.EventsHandler
	productHandler    handler2.ProductHandler
	logger            *logger.Logger
	googleAuthHandler handler3.AuthHandler
}

func NewRouter(
	conf *config.Config,
	eventHandler handler.EventsHandler,
	logger *logger.Logger,
	productHandler handler2.ProductHandler,
	googleAuthHandler handler3.AuthHandler,
) Router {

	router := gin.New()

	return &routerImpl{
		conf:              conf,
		eventHandler:      eventHandler,
		engine:            router,
		logger:            logger,
		productHandler:    productHandler,
		googleAuthHandler: googleAuthHandler,
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

	// Init Google Login
	r.engine.Use(handler3.InitGoogleAuthConnection(r.conf))

	// Required for HTML templates rendering
	//r.engine.LoadHTMLGlob(fmt.Sprintf("%s/templates/*.html", "/"))

	r.engine.GET("/", func(c *gin.Context) {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(c.Writer, false)
	})

	authPath := r.engine.Group("auth")
	r.addAuthPaths(authPath)

	apiV1 := r.engine.Group("v1")
	r.addEventEndpoints(apiV1)
	r.addProductEndpoints(apiV1)

	r.engine.Use(middleware.RecoveryWithZap(r.logger.Desugar(), true))

}

func (r *routerImpl) addAuthPaths(ap *gin.RouterGroup) {

	ap.GET(":provider", r.googleAuthHandler.BeginAuthHandler)
	ap.GET(":provider/callback", r.googleAuthHandler.HandleCallback)

}
