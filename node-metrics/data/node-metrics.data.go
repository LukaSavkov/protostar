package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"metrics-api/domain"
	"metrics-api/errors"
	"net/http"
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

func (nr *NodeMetricsData) ReadMetricsAfterTimestamp(timestamp int64, nodeID string) (json.RawMessage, *errors.ErrorStruct) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

	url := fmt.Sprintf("http://prometheus_healthcheck:9090/api/v1/query_range?query={nodeID=~\"%s\"}&start=%d&end=%d&step=%s",
		nodeID, from, now, step)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
			promResp.Data.Result[i].Values = result.Values[:1] // Keep only the first element
		}
	}

	filteredData, err := json.Marshal(promResp.Data.Result) // Change this to marshal only the results with adjusted values
	if err != nil {
		return nil, errors.NewError("Failed to read response body: "+err.Error(), 500)
	}

	return filteredData, nil
}

func (nr *NodeMetricsData) ReadLastNodeDataWritten(nodeID string) (json.RawMessage, *errors.ErrorStruct) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	currentTime := time.Now().Unix()
	fifteenMinutesAgo := time.Now().Add(-120 * time.Minute).Unix()

	step := calculateStep(fifteenMinutesAgo, currentTime)

	url := fmt.Sprintf("http://prometheus_healthcheck:9090/api/v1/query_range?query={nodeID=~\"%s\"}&start=%d&end=%d&step=%s",
		nodeID, fifteenMinutesAgo, currentTime, step)

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

	var promResp domain.PrometheusResponse
	if err := json.Unmarshal(body, &promResp); err != nil {
		return nil, errors.NewError("Failed to unmarshal response body: "+err.Error(), 500)
	}

	var filteredResults []domain.PrometheusResult
	for _, result := range promResp.Data.Result {
		if metricName, ok := result.Metric["__name__"]; ok && strings.HasPrefix(metricName, "custom_") {
			if len(result.Values) > 0 {
				result.Values = result.Values[:1] // Keep only the first element
			}
			filteredResults = append(filteredResults, result)
		}
	}

	filteredData, err := json.Marshal(filteredResults)
	if err != nil {
		return nil, errors.NewError("Failed to marshal filtered results: "+err.Error(), 500)
	}

	return filteredData, nil
}
