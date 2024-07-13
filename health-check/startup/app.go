package startup

import (
	"context"
	"errors"
	"fmt"
	"health-check/collector"
	"health-check/config"
	"health-check/service"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
)

type app struct {
	config                    *config.AppConfig
	nodeService               *service.NodeService
	shutdownProcesses         []func()
	gracefulShutdownProcesses []func(wg *sync.WaitGroup)
	prometheusService         *service.PrometheusService
	nodes                     *config.NodeConfig
	prometheusRegistry        *prometheus.Registry
}

func NewAppWithConfig(config *config.AppConfig) (*app, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}
	return &app{
		config:                    config,
		shutdownProcesses:         make([]func(), 0),
		gracefulShutdownProcesses: make([]func(wg *sync.WaitGroup), 0),
	}, nil
}

func (a *app) Start() error {
	a.init()
	err := a.startPrometheusServices(a.prometheusRegistry)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (a *app) init() {
	prometheusRegistry := prometheus.NewRegistry()
	a.prometheusRegistry = prometheusRegistry
	a.nodes = config.NewNodeConfig()
	customCollector := collector.NewCustomCollector()
	a.prometheusRegistry.MustRegister(customCollector)
	log.Println(a.config.GetNatsAddress())
	natsConn, err := NewNatsConn(a.config.GetNatsAddress())
	if err != nil {
		log.Fatal(err)
	}
	a.shutdownProcesses = append(a.shutdownProcesses, func() {
		log.Println("closing nats connection")
		natsConn.Close()
	})

	log.Println("Creating Magnetar client...")
	magnetarClient, err := newMagnetarClient(a.config.GetMagnetarAddress())
	if err != nil {
		log.Fatalln("Failed to create magnetar client:", err)
	}

	a.nodeService = service.NewNodeService(magnetarClient, a.nodes)
	a.prometheusService = service.NewPrometheusService(natsConn, a.nodes, a.prometheusRegistry, customCollector)
	log.Println("Scheduling cron jobs")
	a.nodeService.SaveNodes()

	c := cron.New()

	_, err = c.AddFunc("@every 45s", func() {
		log.Println("Executing cron job to save nodes")
		a.nodeService.SaveNodes()
	})
	if err != nil {
		log.Fatalln("Failed to schedule cron job for saving nodes:", err)
	}
	c.Start()

	log.Println("Cron jobs scheduled, entering blocking call")

}

func (a *app) startPrometheusServices(registry *prometheus.Registry) error {
	a.prometheusService.ScheduleNatsRequest()
	promHttpHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	router := mux.NewRouter()
	router.Path("/metrics").Handler(promHttpHandler)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server listening on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
	select {}
}

func (a *app) GracefulStop(ctx context.Context) {
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

func (a *app) shutdown() {
	for _, shutdownProcess := range a.shutdownProcesses {
		shutdownProcess()
	}
}
