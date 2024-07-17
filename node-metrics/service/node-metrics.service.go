package service

import (
	"encoding/json"
	"fmt"
	"metrics-api/data"
	"metrics-api/errors"
)

type NodeMetricsService struct {
	NodeMetricsData *data.NodeMetricsData
}

func NewNodeMetricsService(nodeMetricsData *data.NodeMetricsData) (*NodeMetricsService, error) {
	return &NodeMetricsService{
		NodeMetricsData: nodeMetricsData,
	}, nil
}

func (nm NodeMetricsService) ReadMetricsAfterTimestamp(timestamp int64) (json.RawMessage, *errors.ErrorStruct) {
	fmt.Println("USLO U SERVICE")
	readMetrics, err := nm.NodeMetricsData.ReadMetricsAfterTimestamp(timestamp)
	if err != nil {
		return nil, errors.NewError(err.GetErrorMessage(), 500)
	}
	return readMetrics, nil
}

func (nm NodeMetricsService) ReadMetricsInRange(timestamp, end int64) (json.RawMessage, *errors.ErrorStruct) {
	fmt.Println("USLO U SERVICE")
	readMetrics, err := nm.NodeMetricsData.ReadMetricsInRange(timestamp, end)
	if err != nil {
		return nil, errors.NewError(err.GetErrorMessage(), 500)
	}
	return readMetrics, nil
}

func (nm NodeMetricsService) ReadAppMetrics(app, nodeID string) (json.RawMessage, *errors.ErrorStruct) {
	fmt.Println("USLO U SERVICE")
	readMetrics, err := nm.NodeMetricsData.ReadAppMetrics(app, nodeID)
	if err != nil {
		return nil, errors.NewError(err.GetErrorMessage(), 500)
	}
	return readMetrics, nil
}

func (nm NodeMetricsService) ReadContainerMetrics(container, nodeID string) (json.RawMessage, *errors.ErrorStruct) {
	fmt.Println("USLO U SERVICE")
	readMetrics, err := nm.NodeMetricsData.ReadContainerMetrics(container, nodeID)
	if err != nil {
		return nil, errors.NewError(err.GetErrorMessage(), 500)
	}
	return readMetrics, nil
}

func (nm NodeMetricsService) LastDataWritten(nodeID string) (json.RawMessage, *errors.ErrorStruct) {
	fmt.Println("USLO U SERVICE")
	readMetrics, err := nm.NodeMetricsData.LastDataWritten(nodeID)
	if err != nil {
		return nil, errors.NewError(err.GetErrorMessage(), 500)
	}
	return readMetrics, nil

}

func (nm NodeMetricsService) LastNodeDataWritten(nodeID string) (json.RawMessage, *errors.ErrorStruct) {
	fmt.Println("USLO U SERVICE")
	readMetrics, err := nm.NodeMetricsData.ReadLastNodeDataWritten(nodeID)
	if err != nil {
		return nil, errors.NewError(err.GetErrorMessage(), 500)
	}
	return readMetrics, nil
}

func (nm NodeMetricsService) LastClusterDataWritten(clusterID string) (json.RawMessage, *errors.ErrorStruct) {
	readMetrics, err := nm.NodeMetricsData.ReadLastClusterDataWritten(clusterID)
	if err != nil {
		return nil, errors.NewError(err.GetErrorMessage(), 500)
	}
	return readMetrics, nil
}
