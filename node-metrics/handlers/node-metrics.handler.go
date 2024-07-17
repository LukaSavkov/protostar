package handlers

import (
	"fmt"
	"metrics-api/errors"
	"metrics-api/service"
	"metrics-api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type NodeMetricsHandler struct {
	NodeMetricsService *service.NodeMetricsService
}

func NewNodeMetricsHandler(nodeMetricsService *service.NodeMetricsService) (*NodeMetricsHandler, error) {
	return &NodeMetricsHandler{
		NodeMetricsService: nodeMetricsService,
	}, nil
}

func (nh NodeMetricsHandler) ReadMetricsAfterTimestamp(rw http.ResponseWriter, h *http.Request) {
	fmt.Println("USLO U HANDLER")
	vars := mux.Vars(h)
	timestampStr := vars["timestamp"]
	nodeID := vars["nodeID"]
	timestamp, err1 := strconv.ParseInt(timestampStr, 10, 64)
	fmt.Println("TIMESTAMP JE", timestamp)
	if err1 != nil {
		errors.NewError(err1.Error(), 500)
	}
	returnedMetrics, err := nh.NodeMetricsService.ReadMetricsAfterTimestamp(timestamp, nodeID)
	if err != nil {

		utils.WriteErrorResp(err.GetErrorMessage(), 500, "api/metrics-api/id", rw)
		return
	}
	utils.WriteResp(returnedMetrics, 201, rw)

}

func (nh NodeMetricsHandler) LastDataWritten(rw http.ResponseWriter, h *http.Request) {
	fmt.Println("USLO U HANDLER")
	vars := mux.Vars(h)
	nodeID := vars["nodeID"]

	returnedMetrics, err := nh.NodeMetricsService.LastDataWritten(nodeID)
	if err != nil {

		utils.WriteErrorResp(err.GetErrorMessage(), 500, "api/metrics-api/id", rw)
		return
	}
	utils.WriteResp(returnedMetrics, 201, rw)
}

func (nh NodeMetricsHandler) LastNodeDataWritten(rw http.ResponseWriter, h *http.Request) {
	fmt.Println("USLO U HANDLER")
	vars := mux.Vars(h)
	nodeID := vars["nodeID"]

	returnedMetrics, err := nh.NodeMetricsService.LastNodeDataWritten(nodeID)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), 500, "api/metrics-api/id", rw)
		return
	}
	utils.WriteResp(returnedMetrics, 201, rw)
}

func (nh NodeMetricsHandler) LastClusterDataWritten(rw http.ResponseWriter, h *http.Request) {
	fmt.Println("USLO U HANDLER")
	vars := mux.Vars(h)
	clusterID := vars["clusterID"]

	returnedMetrics, err := nh.NodeMetricsService.LastClusterDataWritten(clusterID)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), 500, "api/metrics-api/id", rw)
		return
	}
	utils.WriteResp(returnedMetrics, 201, rw)
}

func (nh NodeMetricsHandler) Ping(rw http.ResponseWriter, h *http.Request) {
	utils.WriteResp("OK", 200, rw)
}
