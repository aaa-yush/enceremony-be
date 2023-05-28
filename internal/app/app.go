package app

import (
	"context"
	"enceremony-be/internal/router"
	"log"
	"sync"
)

type App interface {
	Start()
}

func NewEnceremonyApp(
	router router.Router) App {
	return &Impl{
		router: router,
	}
}

type Impl struct {
	router router.Router
}

func (i *Impl) Start() {
	var wg sync.WaitGroup

	i.startServer()

	wg.Wait()
}

func (i *Impl) startServer() {

	i.router.MapRoutes()
	err := i.router.Start(context.Background())
	if err != nil {
		log.Fatalf("Couldn't start the app. Got this error: %v", err)
	}
}
