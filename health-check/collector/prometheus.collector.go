package collector

import (
	"health-check/domain"
	"log"
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type CustomCollector struct {
	metrics []domain.MetricData
	mu      sync.Mutex
}

func NewCustomCollector() *CustomCollector {
	return &CustomCollector{}
}

func (collector *CustomCollector) Describe(ch chan<- *prometheus.Desc) {
	// Since our metrics are dynamic, we cannot describe them beforehand
	// and we will leave this method empty.
}

func (collector *CustomCollector) UpdateMetrics(newMetrics []domain.MetricData) {
	collector.mu.Lock()

	log.Println("started updating metrics list")

	collector.metrics = append(collector.metrics, newMetrics...)
	log.Println(len(collector.metrics))

	log.Println("finished updating metrics list")

	collector.mu.Unlock()
}

func (collector *CustomCollector) Collect(ch chan<- prometheus.Metric) {
	collector.mu.Lock()

	log.Println("started collecting metrics")

	seenMetrics := make(map[string]bool)

	for _, metricData := range collector.metrics {
		// wtf
		if strings.HasPrefix(metricData.MetricName, "go_") || strings.HasPrefix(metricData.MetricName, "process_") {
			continue
		}
		labels := make([]string, 0, len(metricData.Labels))
		labelValues := make([]string, 0, len(metricData.Labels))
		for key, value := range metricData.Labels {
			labels = append(labels, key)
			labelValues = append(labelValues, value)
		}

		metricKey := metricData.MetricName + "-" + joinLabels(labels, labelValues)
		if !seenMetrics[metricKey] {
			seenMetrics[metricKey] = true
			desc := prometheus.NewDesc(
				metricData.MetricName,
				"Custom metric collected from external source",
				labels, nil,
			)
			ch <- prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				metricData.Value,
				labelValues...,
			)
		} else {
			log.Println(metricKey)
		}
	}
	log.Println(len(collector.metrics))
	collector.metrics = make([]domain.MetricData, 0)
	log.Println(len(collector.metrics))

	log.Println("finished collecting metrics")

	collector.mu.Unlock()
}

func joinLabels(labelKeys, labelValues []string) string {
	result := ""
	for _, value := range labelKeys {
		result += value + "-"
	}
	for _, value := range labelValues {
		result += value + "-"
	}
	return result
}
