package collector

import (
	"health-check/domain"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type CustomCollector struct {
	metrics []domain.MetricData
	mu      *sync.Mutex
}

func NewCustomCollector() *CustomCollector {
	return &CustomCollector{
		mu: &sync.Mutex{},
	}
}

func (collector *CustomCollector) Describe(ch chan<- *prometheus.Desc) {
	// Since our metrics are dynamic, we cannot describe them beforehand
	// and we will leave this method empty.
}

func (collector *CustomCollector) UpdateMetrics(newMetrics []domain.MetricData) {
	collector.mu.Lock()
	defer collector.mu.Unlock()
	collector.metrics = append(collector.metrics, newMetrics...)
}

func (collector *CustomCollector) Collect(ch chan<- prometheus.Metric) {
	collector.mu.Lock()
	defer collector.mu.Unlock()

	seenMetrics := make(map[string]bool)

	for _, metricData := range collector.metrics {
		labels := make([]string, 0, len(metricData.Labels))
		labelValues := make([]string, 0, len(metricData.Labels))
		for key, value := range metricData.Labels {
			labels = append(labels, key)
			labelValues = append(labelValues, value)
		}

		desc := prometheus.NewDesc(
			metricData.MetricName,
			"Custom metric collected from external source",
			labels, nil,
		)

		metricKey := metricData.MetricName + "-" + joinLabels(labelValues)
		if !seenMetrics[metricKey] {
			ch <- prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				metricData.Value,
				labelValues...,
			)
			seenMetrics[metricKey] = true
		}
	}
	collector.metrics = make([]domain.MetricData, 0)
}

func joinLabels(labelValues []string) string {
	result := ""
	for _, value := range labelValues {
		result += value + "-"
	}
	return result
}
