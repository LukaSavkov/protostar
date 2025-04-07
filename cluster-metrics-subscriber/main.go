package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/nats-io/nats.go"
)

type ClusterMetric struct {
	Name  string
	Value float64
}

func (c *ClusterMetric) UnmarshalJSON(data []byte) error {
	metricAny := make(map[string]any)
	err := json.Unmarshal(data, &metricAny)
	if err != nil {

		return err
	}
	metricNameAny, ok := metricAny["metric"].(map[string]any)
	if !ok {
		c.Name = ""
	} else {
		c.Name = metricNameAny["__name__"].(string)
	}
	valuesAny, ok := metricAny["value"].([]any)
	if !ok || len(valuesAny) < 2 {
		c.Value = 0
	} else {
		valueStr := valuesAny[1].(string)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			log.Println(err)
			c.Value = 0
		} else {
			c.Value = value
		}
	}
	return nil
}

type ClusterMetricsMsg struct {
	TimestampNano int64           `json:"timestamp"`
	ClusterId     string          `json:"clusterId"`
	Metrics       []ClusterMetric `json:"metrics"`
}

func main() {
	conn, err := nats.Connect(fmt.Sprintf("nats://%s", "localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	conn.Subscribe("metrics.clusters.*", func(msg *nats.Msg) {
		metrics := new(ClusterMetricsMsg)
		err := json.Unmarshal(msg.Data, metrics)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(metrics)
	})
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, os.Kill)

	<-stopChan
	log.Println("Received termination signal, shutdown ...")
}
