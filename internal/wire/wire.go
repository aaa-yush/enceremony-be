//go:build wireinject
// +build wireinject

package wire

import (
	mysql2 "enceremony-be/commons/clients/mysql"
	"enceremony-be/internal/app"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"enceremony-be/internal/database/mysql"
	"enceremony-be/internal/events/handler"
	"enceremony-be/internal/events/repo"
	"enceremony-be/internal/events/service"
	handler2 "enceremony-be/internal/product/handler"
	repo2 "enceremony-be/internal/product/repo"
	service2 "enceremony-be/internal/product/service"
	"enceremony-be/internal/router"
	"github.com/google/wire"
)

func InitializeApp() (app.App, error) {

	panic(
		wire.Build(
			config.NewConfig,
			repo.NewEventsRepo,
			handler.NewEventsHandler,
			service.NewEventService,
			router.NewRouter,

			config.NewLoggerConf,
			logger.NewLogger,
			mysql.NewMysqlStore,
			mysql2.NewMysqlConnection,
			config.NewMysqlConf,
			handler2.NewProductHandler,
			service2.NewProductService,
			repo2.NewProductRepo,

			app.NewEnceremonyApp,
		))
}
