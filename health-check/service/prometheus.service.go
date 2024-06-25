package service

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/robfig/cron/v3"
	"health-check/collector"
	"health-check/config"
	"health-check/domain"
	"log"
	"strings"
	"sync"
	"time"
)

type PrometheusService struct {
	regMetrics         map[string]*prometheus.GaugeVec
	ns                 *nats.Conn
	nodes              *config.NodeConfig
	prometheusRegistry *prometheus.Registry
	mutex              sync.Mutex
	collector          *collector.CustomCollector
}

func NewPrometheusService(ns *nats.Conn, nodes *config.NodeConfig, prometheusRegistry *prometheus.Registry, collector *collector.CustomCollector) *PrometheusService {
	return &PrometheusService{
		regMetrics:         make(map[string]*prometheus.GaugeVec),
		ns:                 ns,
		nodes:              nodes,
		prometheusRegistry: prometheusRegistry,
		collector:          collector,
	}
}

//func (ps *PrometheusService) RegisterMetrics(metrics domain.MetricFileFormat) {
//
//	foundNode := ps.nodes.GetNode(metrics.NodeID)
//
//	for _, metric := range metrics.Metrics {
//		metric.Labels["nodeID"] = metrics.NodeID
//		metricNameID := strings.ReplaceAll(uuid.New().String(), "-", "_")
//		metricKey, labelSlice, labelValues := GenerateRegisterKey(metric)
//		if vec, exists := ps.regMetrics[metricKey]; !exists {
//			vec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
//				Name: metric.MetricName + "_" + metricNameID,
//				Help: "",
//			}, labelSlice)
//			ps.prometheusRegistry.MustRegister(vec)
//			ps.regMetrics[metricKey] = vec
//			vec.With(labelValues).Set(metric.Value)
//		} else {
//			vec.With(labelValues).Set(metric.Value)
//		}
//
//		service, exists := metric.Labels["container_label_com_docker_compose_service"]
//		if exists {
//			if _, doesServiceAlreadyExist := foundNode.Services[service]; doesServiceAlreadyExist {
//				continue
//			}
//			foundNode.Services[service] = true
//		}
//	}
//
//	foundNode.LastSeen = time.Now()
//	err := ps.publishToNATS(&foundNode)
//	if err != nil {
//		log.Println("Error publishing to NATS:", err)
//	}
//
//}
//func GenerateRegisterKey(metric domain.MetricData) (string, []string, prometheus.Labels) {
//	labelSlice := make([]string, 0, len(metric.Labels))
//	labelValues := prometheus.Labels{}
//
//	for k := range metric.Labels {
//		labelSlice = append(labelSlice, k)
//		labelValues[k] = metric.Labels[k]
//	}
//
//	sort.Strings(labelSlice)
//	labelKey := strings.Join(labelSlice, "_")
//	metricKey := fmt.Sprintf("%s_%s", metric.MetricName, labelKey)
//
//	return metricKey, labelSlice, labelValues
//}

func (ps *PrometheusService) ScheduleNatsRequest(natsSubject string) {
	c := cron.New()
	_, err := c.AddFunc("@every 60s", func() {
		ps.HandleNatsRequest(natsSubject)
	})
	if err != nil {
		log.Println("Error scheduling cron job:", err)
	}
	c.Start()
}

func (ps *PrometheusService) HandleNatsRequest(natsSubject string) {
	log.Println("USLO U NATS")
	response, err := ps.ns.Request(natsSubject, []byte("metrics"), 10*time.Second)
	if err != nil {
		log.Println("Error making request:", err)
		return
	}

	var metrics domain.MetricFileFormat
	err = json.Unmarshal(response.Data, &metrics)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	log.Println(metrics)

	for _, metric := range metrics.Metrics {
		metric.Labels["nodeID"] = metrics.NodeID
		foundNode := ps.nodes.GetNode(metrics.NodeID)
		if foundNode.Services == nil {
			foundNode.Services = make(map[string][]domain.MetricData)
		}
		service, exists := metric.Labels["container_label_com_docker_compose_service"]
		if exists {
			if _, doesServiceAlreadyExist := foundNode.Services[service]; !doesServiceAlreadyExist {
				foundNode.Services[service] = make([]domain.MetricData, 0)
			}
		}

		if strings.Contains(metric.MetricName, "custom_service") {

			foundNode.Services[service] = append(foundNode.Services[service], metric)

		}

		foundNode.LastSeen = time.Now()
	}
	log.Println(metrics.Metrics)
	err = ps.PublishNodesToNATS(ps.nodes)
	if err != nil {
		log.Println("Error publishing to NATS:", err)
	}
	ps.collector.UpdateMetrics(metrics.Metrics)
}

func (ps *PrometheusService) PublishNodesToNATS(data *config.NodeConfig) error {
	msg, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ps.ns.Publish("howAreYou?", msg)
	if err != nil {
		return err
	}
	log.Println("Published data to NATS:", data)
	return nil
}
