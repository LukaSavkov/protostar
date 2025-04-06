package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"time"

	magnetarapi "github.com/c12s/magnetar/pkg/api"
	"github.com/nats-io/nats.go"
)

type ClusterMetricsService struct {
	nodeMetrics *NodeMetricsService
	publisher   *nats.Conn
	magnetar    magnetarapi.MagnetarClient
}

func NewClusterMetricsService(nodeMetrics *NodeMetricsService, publisher *nats.Conn, magnetar magnetarapi.MagnetarClient) (*ClusterMetricsService, error) {
	return &ClusterMetricsService{
		nodeMetrics: nodeMetrics,
		publisher:   publisher,
		magnetar:    magnetar,
	}, nil
}

func (s *ClusterMetricsService) Publish() {
	for _, cluster := range s.listClusters() {
		log.Println("cluster " + cluster)
		metrics, err := s.nodeMetrics.LastClusterDataWritten(cluster)
		if err != nil {
			log.Println(err.GetErrorStatus())
			log.Println(err.GetErrorMessage())
			continue
		}
		msg := map[string]any{
			"clusterId": cluster,
			"timestamp": time.Now().UnixNano(),
			"metrics":   metrics,
		}
		natsMsg, err2 := json.Marshal(msg)
		if err2 != nil {
			log.Println(err2)
			continue
		}
		log.Println(string(natsMsg))
		err2 = s.publisher.Publish(fmt.Sprintf("metrics.clusters.%s", cluster), natsMsg)
		if err2 != nil {
			log.Println(err2)
			continue
		}
	}
}

func (s *ClusterMetricsService) listClusters() []string {
	clusters := []string{}
	resp, err := s.magnetar.ListAllNodes(context.TODO(), &magnetarapi.ListAllNodesReq{})
	if err != nil {
		log.Println(err)
		return clusters
	}
	for _, node := range resp.Nodes {
		if node.Org != "" && !slices.Contains(clusters, node.Org) {
			clusters = append(clusters, node.Org)
		}
	}
	return clusters
}
