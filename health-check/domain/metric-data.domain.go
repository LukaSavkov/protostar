package domain

import (
	"encoding/json"
	"health-check/errors"
	"io"
	"time"
)

type MetricData struct {
	MetricName string            `json:"metric_name"`
	Labels     map[string]string `json:"labels"`
	Value      float64           `json:"value"`
	Timestamp  int64             `json:"timestamp"`
}
type MetricFileFormat struct {
	Metrics   []MetricData `json:"metrics"`
	NodeID    string       `json:"nodeId"`
	ClusterId string       `json:"clusterId"`
}
type MetricsSendingDataHolder struct {
	Services map[string]bool
	Metrics  []MetricData
}
type MetricsSending struct {
	Data map[string]MetricsSendingDataHolder
}

func (md MetricData) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(md)
}
func (md MetricData) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(md)
}

type Node struct {
	NodeID   string                  `json:"nodeID"`
	Services map[string][]MetricData `json:"services"`
	LastSeen time.Time               `json:"lastSeen"`
}

type NodeIDs struct {
	IDs []string `json:"ids"`
}

type NodeRegistryData interface {
	WriteMetrics(md MetricData) (*MetricData, *errors.ErrorStruct)
}

func NewNode(nodeID string) *Node {
	return &Node{
		NodeID:   nodeID,
		Services: make(map[string][]MetricData),
	}
}
