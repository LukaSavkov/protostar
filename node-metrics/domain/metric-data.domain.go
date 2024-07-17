package domain

import (
	"encoding/json"
	"io"
)

type MetricData struct {
	MetricName string            `json:"metric_name"`
	Labels     map[string]string `json:"labels"`
	Value      float64           `json:"value"`
	Timestamp  int64             `json:"timestamp"`
}

type MetricFileFormat struct {
	Metrics []MetricData `json:"metrics"`
	NodeID  string       `json:"nodeID"`
}

type PrometheusResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string             `json:"resultType"`
		Result     []PrometheusResult `json:"result"`
	} `json:"data"`
}

type PrometheusSingleResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string                   `json:"resultType"`
		Result     []PrometheusSingleResult `json:"result"`
	} `json:"data"`
}

type PrometheusResult struct {
	Metric map[string]string `json:"metric"`
	Values [][]interface{}   `json:"values"`
}

type PrometheusSingleResult struct {
	Metric map[string]string `json:"metric"`
	Values []interface{}     `json:"value"`
}

type FilteredMetric struct {
	Metric     json.RawMessage `json:"metric"`
	FirstValue interface{}     `json:"first_value"`
}

func (md MetricData) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(md)
}

func (md MetricData) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(md)
}
