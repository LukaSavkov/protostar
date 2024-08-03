package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"metrics-api/domain"
	"metrics-api/errors"
	"net/http"
	"slices"
	"strings"
	"time"
)

type NodeMetricsData struct {
	client *http.Client
}

func NewMetricRepo(client *http.Client) (*NodeMetricsData, error) {
	return &NodeMetricsData{
		client: client,
	}, nil
}

func calculateStep(start, end int64) string {
	maxDataPoints := 10000
	step := (end - start) / int64(maxDataPoints)
	if step < 15 {
		step = 15
	}
	return fmt.Sprintf("%ds", step)
}

func (nr *NodeMetricsData) ReadMetricsAfterTimestamp(timestamp int64) (json.RawMessage, *errors.ErrorStruct) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("Entering ReadMetricsAfterTimestamp")
	fmt.Println("Timestamp is", timestamp)

	now := time.Now().Unix()
	from := timestamp

	fmt.Println("Timestamp now is", now)
	fmt.Println("Timestamp from is", from)

	// Calculate step to aim for no more than 10000 points
	step := calculateStep(timestamp, time.Now().Unix())
	fmt.Printf("Using step size of %s seconds\n", step)

	url := fmt.Sprintf("http://prometheus_healthcheck:9090/api/v1/query_range?query={__name__!=\"\"}&start=%d&end=%d&step=%s", from, now, step)
	fmt.Println("URL is", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return nil, errors.NewError("Failed to create HTTP request: "+err.Error(), 500)
	}

	resp, err := nr.client.Do(req)
	if err != nil {
		fmt.Println("Error during request")
		return nil, errors.NewError("HTTP request failed: "+err.Error(), 500)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected HTTP status")
		return nil, errors.NewError(fmt.Sprintf("Unexpected HTTP status: %d", resp.StatusCode), resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}
	fmt.Println("Raw JSON Body", string(body))

	return body, nil
}

func (nr *NodeMetricsData) ReadMetricsInRange(timestamp, end int64) (json.RawMessage, *errors.ErrorStruct) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("Entering ReadMetricsAfterTimestamp")
	fmt.Println("Timestamp is", timestamp)

	from := timestamp

	fmt.Println("Timestamp now is", end)
	fmt.Println("Timestamp from is", from)

	// Calculate step to aim for no more than 10000 points
	step := calculateStep(timestamp, end)
	fmt.Printf("Using step size of %s seconds\n", step)

	url := fmt.Sprintf("http://prometheus_healthcheck:9090/api/v1/query_range?query={__name__!=\"\"}&start=%d&end=%d&step=%s", from, end, step)
	fmt.Println("URL is", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return nil, errors.NewError("Failed to create HTTP request: "+err.Error(), 500)
	}

	resp, err := nr.client.Do(req)
	if err != nil {
		fmt.Println("Error during request")
		return nil, errors.NewError("HTTP request failed: "+err.Error(), 500)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected HTTP status")
		return nil, errors.NewError(fmt.Sprintf("Unexpected HTTP status: %d", resp.StatusCode), resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}
	fmt.Println("Raw JSON Body", string(body))

	return body, nil
}

func (nr *NodeMetricsData) ReadAppMetrics(app, nodeID string) (json.RawMessage, *errors.ErrorStruct) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	start := time.Now().Add(-12 * time.Hour).Unix()
	// Calculate step to aim for no more than 10000 points
	step := calculateStep(start, time.Now().Unix())
	fmt.Printf("Using step size of %s seconds\n", step)

	url := fmt.Sprintf("http://prometheus_healthcheck:9090/api/v1/query_range?query={app=\"%s\",nodeID=\"%s\"}&start=%d&end=%d&step=%s", app, nodeID, start, time.Now().Unix(), step)
	fmt.Println("URL is", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return nil, errors.NewError("Failed to create HTTP request: "+err.Error(), 500)
	}

	resp, err := nr.client.Do(req)
	if err != nil {
		fmt.Println("Error during request")
		return nil, errors.NewError("HTTP request failed: "+err.Error(), 500)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected HTTP status")
		return nil, errors.NewError(fmt.Sprintf("Unexpected HTTP status: %d", resp.StatusCode), resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}
	fmt.Println("Raw JSON Body", string(body))

	return body, nil
}

func (nr *NodeMetricsData) ReadContainerMetrics(container, nodeID string) (json.RawMessage, *errors.ErrorStruct) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	start := time.Now().Add(-12 * time.Hour).Unix()
	// Calculate step to aim for no more than 10000 points
	step := calculateStep(start, time.Now().Unix())
	fmt.Printf("Using step size of %s seconds\n", step)

	url := fmt.Sprintf("http://prometheus_healthcheck:9090/api/v1/query_range?query={image!=\"\",name=\"%s\",nodeID=\"%s\"}&start=%d&end=%d&step=%s", container, nodeID, start, time.Now().Unix(), step)
	fmt.Println("URL is", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return nil, errors.NewError("Failed to create HTTP request: "+err.Error(), 500)
	}

	resp, err := nr.client.Do(req)
	if err != nil {
		fmt.Println("Error during request")
		return nil, errors.NewError("HTTP request failed: "+err.Error(), 500)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected HTTP status")
		return nil, errors.NewError(fmt.Sprintf("Unexpected HTTP status: %d", resp.StatusCode), resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}
	fmt.Println("Raw JSON Body", string(body))

	return body, nil
}

func (nr *NodeMetricsData) LastDataWritten(nodeID string) (json.RawMessage, *errors.ErrorStruct) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	currentTime := time.Now().Unix()
	fifteenMinutesAgo := time.Now().Add(-120 * time.Minute).Unix()

	step := calculateStep(fifteenMinutesAgo, currentTime)

	url := fmt.Sprintf("http://prometheus_healthcheck:9090/api/v1/query_range?query={nodeID=~\"%s\"}&start=%d&end=%d&step=%s",
		nodeID, fifteenMinutesAgo, currentTime, step)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}

	resp, err := nr.client.Do(req)
	if err != nil {
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}

	var promResp domain.PrometheusResponse
	if err := json.Unmarshal(body, &promResp); err != nil {
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}

	// Only retain the first value in the values array for each metric
	for i, result := range promResp.Data.Result {
		if len(result.Values) > 0 {
			promResp.Data.Result[i].Values = result.Values[len(result.Values)-1:] // Keep only the first element
		}
	}

	filteredData, err := json.Marshal(promResp.Data.Result) // Change this to marshal only the results with adjusted values
	if err != nil {
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}

	return filteredData, nil
}

func (nr *NodeMetricsData) ReadLastNodeDataWritten(nodeID string) (json.RawMessage, *errors.ErrorStruct) {
	urlPrefix := "http://prometheus_healthcheck:9090/api/v1/query?query="

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	url := fmt.Sprintf("http://prometheus_healthcheck:9090/api/v1/query?query=last_over_time({nodeID=~\"%s\"}[10m])", nodeID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, errors.NewError("Failed to create HTTP request: "+err.Error(), 500)
	}

	resp, err := nr.client.Do(req)
	if err != nil {
		return nil, errors.NewError("HTTP request failed: "+err.Error(), 500)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewError(fmt.Sprintf("Unexpected HTTP status: %d", resp.StatusCode), resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}

	var promResp domain.PrometheusSingleResponse
	if err := json.Unmarshal(body, &promResp); err != nil {
		return nil, errors.NewError("Failed to unmarshal response body: "+err.Error(), 500)
	}

	var filteredResults []domain.PrometheusSingleResult
	for _, result := range promResp.Data.Result {
		if metricName, ok := result.Metric["__name__"]; ok && strings.HasPrefix(metricName, "custom_") && !slices.Contains(custom_calculated_metrics, metricName) {
			filteredResults = append(filteredResults, result)
		}
	}

	url = fmt.Sprintf("%s100*(scalar(last_over_time(machine_cpu_cores{nodeID=~\"%s\"}[1m]))-sum(avg by (cpu)(rate(node_cpu_seconds_total{mode=\"idle\",nodeID=\"%s\"}[1m]))))", urlPrefix, nodeID, nodeID)
	url = strings.ReplaceAll(url, " ", "%20")
	log.Println(url)
	promResp2, promErr := nr.query(url)
	if promErr != nil {
		return nil, promErr
	}
	for _, result := range promResp2.Data.Result {
		result.Metric["__name__"] = "custom_node_cpu_usage_percentage"
		filteredResults = append(filteredResults, result)
	}

	url = fmt.Sprintf("%srate(container_cpu_usage_seconds_total{nodeID=\"%s\",mode!=\"idle\",image!=\"\"}[1m])*100", urlPrefix, nodeID)
	log.Println(url)
	promResp2, promErr = nr.query(url)
	if promErr != nil {
		return nil, promErr
	}
	for _, result := range promResp2.Data.Result {
		result.Metric["__name__"] = "custom_service_cpu_usage"
		filteredResults = append(filteredResults, result)
	}

	filteredData, err := json.Marshal(filteredResults)
	if err != nil {
		return nil, errors.NewError("Failed to marshal filtered results: "+err.Error(), 500)
	}

	return filteredData, nil
}

func (nr *NodeMetricsData) ReadLastClusterDataWritten(clusterID string) (json.RawMessage, *errors.ErrorStruct) {
	urlPrefix := "http://prometheus_healthcheck:9090/api/v1/query?query="
	url := fmt.Sprintf("%slast_over_time({clusterID=~\"%s\"}[10m])", urlPrefix, clusterID)
	promResp, promErr := nr.query(url)
	if promErr != nil {
		return nil, promErr
	}

	var filteredResults []domain.PrometheusSingleResult
	for _, result := range promResp.Data.Result {
		if metricName, ok := result.Metric["__name__"]; ok && strings.HasPrefix(metricName, "custom_service") && !slices.Contains(custom_calculated_metrics, metricName) {
			filteredResults = append(filteredResults, result)
		}
	}

	urlTemplate := "%ssum(last_over_time(%s{clusterID=~\"%s\"}[10m]))"
	for _, metricName := range custom_node_metrics {
		url := fmt.Sprintf(urlTemplate, urlPrefix, metricName, clusterID)
		promResp, promErr := nr.query(url)
		if promErr != nil {
			return nil, promErr
		}
		for _, result := range promResp.Data.Result {
			result.Metric["__name__"] = metricName
			filteredResults = append(filteredResults, result)
		}
	}

	url = fmt.Sprintf("%s100*(scalar(sum(last_over_time(machine_cpu_cores{clusterID=~\"%s\"}[1m])))-sum(avg by (nodeID,cpu)(rate(node_cpu_seconds_total{mode=\"idle\",clusterID=\"%s\"}[1m]))))", urlPrefix, clusterID, clusterID)
	url = strings.ReplaceAll(url, " ", "%20")
	promResp, promErr = nr.query(url)
	if promErr != nil {
		return nil, promErr
	}
	for _, result := range promResp.Data.Result {
		result.Metric["__name__"] = "custom_node_cpu_usage_percentage"
		filteredResults = append(filteredResults, result)
	}

	url = fmt.Sprintf("%srate(container_cpu_usage_seconds_total{clusterID=\"%s\",mode!=\"idle\",image!=\"\"}[1m])*100", urlPrefix, clusterID)
	promResp, promErr = nr.query(url)
	if promErr != nil {
		return nil, promErr
	}
	for _, result := range promResp.Data.Result {
		result.Metric["__name__"] = "custom_service_cpu_usage"
		filteredResults = append(filteredResults, result)
	}

	filteredData, err := json.Marshal(filteredResults)
	if err != nil {
		return nil, errors.NewError("Failed to marshal filtered results: "+err.Error(), 500)
	}

	return filteredData, nil
}

func (nr *NodeMetricsData) query(url string) (*domain.PrometheusSingleResponse, *errors.ErrorStruct) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, errors.NewError("Failed to create HTTP request: "+err.Error(), 500)
	}

	resp, err := nr.client.Do(req)
	if err != nil {
		return nil, errors.NewError("HTTP request failed: "+err.Error(), 500)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			log.Println(string(body))
		}
		return nil, errors.NewError(fmt.Sprintf("Unexpected HTTP status: %d", resp.StatusCode), resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}

	var promResp *domain.PrometheusSingleResponse = new(domain.PrometheusSingleResponse)
	if err := json.Unmarshal(body, promResp); err != nil {
		return nil, errors.NewError("Failed to unmarshal response body: "+err.Error(), 500)
	}
	return promResp, nil
}

var custom_node_metrics = []string{
	"custom_node_ram_available_mb",
	"custom_node_ram_total_mb",
	"custom_node_disk_usage_gb",
	"custom_node_disk_total_gb",
	"custom_node_network_receive_mb",
	"custom_node_network_transmit_mb",
}

var custom_calculated_metrics = []string{
	"custom_node_cpu_usage_percentage",
	"custom_service_cpu_usage",
}
