package startup

import (
	"context"
	"errors"
	"fmt"
	"log"
	"metrics-api/config"
	"metrics-api/data"
	"metrics-api/handlers"
	"metrics-api/servers"
	"metrics-api/service"
	"net/http"
	"sync"
	"time"
)

type App struct {
	appConfig                 *config.AppConfig
	httpServer                *servers.HttpServer
	shutdownProcesses         []func()
	gracefulShutdownProcesses []func(wg *sync.WaitGroup)
}

func NewAppWithConfig(config *config.AppConfig) (*App, error) {

	if config == nil {
		return nil, errors.New("config is nil")
	}

	app := &App{
		appConfig:                 config,
		shutdownProcesses:         make([]func(), 0),
		gracefulShutdownProcesses: make([]func(wg *sync.WaitGroup), 0),
	}
	app.init()
	return app, nil
}

func (a *App) startHttpServer() {
	a.httpServer.InitServer(a.appConfig.GetServerAddress())
	fmt.Println("Uslo je da kreira server sa adresom", a.appConfig.GetServerAddress())

	a.httpServer.Run()
}

func (a *App) init() {
	fmt.Println("USLO U INIT")

	client := &http.Client{
		Timeout: time.Second * 120, // Set a 30-second timeout for all requests
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}

	nodeMetricsData, err := data.NewMetricRepo(client)
	if err != nil {
		fmt.Println(err.Error())
	}
	metricsService, err := service.NewNodeMetricsService(nodeMetricsData)
	if err != nil {
		fmt.Println(err.Error())
	}
	metricsHandler, err := handlers.NewNodeMetricsHandler(metricsService)
	if err != nil {
		fmt.Println(err.Error())
	}
	customHttpServer := servers.NewHttpServer(metricsHandler)
	a.httpServer = customHttpServer
	a.startHttpServer()

}
func (a *App) GracefulStop(ctx context.Context) {
	// call all shutdown processes after a timeout or graceful shutdown processes completion
	defer a.shutdown()

	// wait for all graceful shutdown processes to complete
	wg := &sync.WaitGroup{}
	wg.Add(len(a.gracefulShutdownProcesses))

	for _, gracefulShutdownProcess := range a.gracefulShutdownProcesses {
		go gracefulShutdownProcess(wg)
	}

	// notify when graceful shutdown processes are done
	gracefulShutdownDone := make(chan struct{})
	go func() {
		wg.Wait()
		gracefulShutdownDone <- struct{}{}
	}()

	// wait for graceful shutdown processes to complete or for ctx timeout
	select {
	case <-ctx.Done():
		log.Println("ctx timeout ... shutting down")
	case <-gracefulShutdownDone:
		log.Println("app gracefully stopped")
	}
}

func (a *App) shutdown() {
	for _, shutdownProcess := range a.shutdownProcesses {
		shutdownProcess()
	}
}
