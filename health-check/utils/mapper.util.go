package utils

import "github.com/prometheus/client_golang/prometheus"

func ConvertFromStringArrayToMap(list []string) map[string]bool {
	result := make(map[string]bool)
	for _, element := range list {
		result[element] = true

	}
	return result
}

func ConvertFromLabelsMapToStringArrayWithPrometheusLabels(labelsMap map[string]bool, labels map[string]string) ([]string, prometheus.Labels) {
	var stringSlice []string = make([]string, 0)
	prometheusLabels := prometheus.Labels{}
	for k, _ := range labelsMap {
		stringSlice = append(stringSlice, k)
		if _, exist := labels[k]; exist {
			prometheusLabels[k] = labels[k]
		} else {
			prometheusLabels[k] = ""
		}
	}
	return stringSlice, prometheusLabels
}
