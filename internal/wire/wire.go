package wire

import (
	"enceremony-be/internal/app"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	"enceremony-be/internal/events/handler"
	"enceremony-be/internal/events/repo"
	"enceremony-be/internal/events/service"
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

			logger.NewLogger,

			app.NewEnceremonyApp,
		))
}
