package config

import (
	"health-check/domain"
)

type NodeConfig struct {
	nodes     map[string]domain.Node
	loadedIDs []string
}

func NewNodeConfig() *NodeConfig {
	return &NodeConfig{
		nodes:     make(map[string]domain.Node),
		loadedIDs: make([]string, 0),
	}
}

func (nc NodeConfig) GetNodes() map[string]domain.Node {
	return nc.nodes
}

func (nc *NodeConfig) GetNode(name string) domain.Node {
	return nc.nodes[name]

}

func (nc *NodeConfig) SetNodes(nodes map[string]domain.Node) {
	nc.nodes = nodes
}

func (nc *NodeConfig) RemoveNode(name string) {
	delete(nc.nodes, name)
}

func (nc *NodeConfig) RemoveNodes() {
	nc.nodes = make(map[string]domain.Node)
}
func (nc *NodeConfig) GetLoadedIDs() []string {
	return nc.loadedIDs
}
func (nc *NodeConfig) AppendLoadedIDs(id string) {
	nc.loadedIDs = append(nc.loadedIDs, id)
}

func (nc *NodeConfig) AppendNewNode(nodeID string) domain.Node {
	createdNode := domain.Node{
		NodeID: nodeID,
	}
	nc.nodes[nodeID] = createdNode
	return createdNode
}
