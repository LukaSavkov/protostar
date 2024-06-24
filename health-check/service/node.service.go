package service

import (
	"context"
	magnetarapi "github.com/c12s/magnetar/pkg/api"
	"health-check/config"
	"health-check/domain"
	"health-check/mappers"
	"log"
	"sync"
	"time"
)

type NodeService struct {
	magnetar magnetarapi.MagnetarClient
	nodeIDs  domain.NodeIDs
	mu       sync.RWMutex
	nodes    *config.NodeConfig
}

func NewNodeService(magnetar magnetarapi.MagnetarClient, nodes *config.NodeConfig) *NodeService {
	return &NodeService{
		magnetar: magnetar,
		nodes:    nodes,
	}
}
func (ns *NodeService) SaveNodes() {
	log.Println("Fetching node pool...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := ns.magnetar.ListNodePool(ctx, &magnetarapi.ListNodePoolReq{})
	if err != nil {
		log.Println("Error fetching node pool:", err)
		return
	}

	mappers.MapFromApiExternalApplicationToModelExternalApplication(resp, ns.nodes)
	//
	//ns.mu.Lock()
	//ns.nodes.SetNodes(*mappedValue)
	//ns.mu.Unlock()

	log.Println("Nodes: ", ns.nodes.GetNodes())
}

func (ns *NodeService) GetNodeIDs() []string {
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	return ns.nodeIDs.IDs
}
