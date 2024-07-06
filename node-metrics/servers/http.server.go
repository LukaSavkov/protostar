package servers

import (
	"context"
	"fmt"
	"log"
	"metrics-api/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	server         *http.Server
	metricsHandler *handlers.NodeMetricsHandler
}

func NewHttpServer(metricsHandler *handlers.NodeMetricsHandler) *HttpServer {
	return &HttpServer{
		metricsHandler: metricsHandler,
	}
}

func (httpServer *HttpServer) ConfigureRouter() *mux.Router {
	fmt.Println("USLO U SERVER CONFIGURE")

	router := mux.NewRouter()
	router.HandleFunc("/api/metrics-api/latest-data/{nodeID}", httpServer.metricsHandler.LastDataWritten).Methods("GET")
	router.HandleFunc("/api/metrics-api/latest-node-data/{nodeID}", httpServer.metricsHandler.LastNodeDataWritten).Methods("GET")
	router.HandleFunc("/api/metrics-api/ping", httpServer.metricsHandler.Ping).Methods("GET")
	router.HandleFunc("/api/metrics-api/{timestamp}/{nodeID}", httpServer.metricsHandler.ReadMetricsAfterTimestamp).Methods("GET")

	return router
}

func (httpServer *HttpServer) InitServer(port string) {
	httpServer.server = &http.Server{
		Addr:         ":" + port,
		Handler:      httpServer.ConfigureRouter(),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
}

func (httpServer *HttpServer) GetHttpServer() *http.Server {
	return httpServer.server
}

func (httpServer *HttpServer) Run() {
	go func() {
		log.Println("HTTP Server running.")
		if err := httpServer.server.ListenAndServe(); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, os.Kill)

	<-stopChan
	log.Println("Received terminate, graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.server.Shutdown(ctx); err != nil {
		log.Fatalf("Cannot gracefully shutdown: %v", err)
	}
	log.Println("Server stopped")
}
