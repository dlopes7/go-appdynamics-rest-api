package appdrest

import (
	"fmt"
)

// Node represents one node within one Application
type Node struct {
	AppAgentVersion     string      `json:"appAgentVersion"`
	MachineAgentVersion string      `json:"machineAgentVersion"`
	AgentType           string      `json:"agentType"`
	Type                string      `json:"type"`
	MachineName         string      `json:"machineName"`
	AppAgentPresent     bool        `json:"appAgentPresent"`
	NodeUniqueLocalID   string      `json:"nodeUniqueLocalId"`
	MachineID           int         `json:"machineId"`
	MachineOSType       string      `json:"machineOSType"`
	TierID              int         `json:"tierId"`
	TierName            string      `json:"tierName"`
	MachineAgentPresent bool        `json:"machineAgentPresent"`
	Name                string      `json:"name"`
	IPAddresses         interface{} `json:"ipAddresses"`
	ID                  int         `json:"id"`
}

// NodeService intermediates Node requests
type NodeService service

// GetNodes obtains all Nodes from an Application
func (s *NodeService) GetNodes(appIDOrName string) ([]*Node, error) {

	url := fmt.Sprintf("controller/rest/applications/%s/nodes?output=json", appIDOrName)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var nodes []*Node
	err = s.client.Do(req, &nodes)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

// GetNode obtains a single Node from an Application
func (s *NodeService) GetNode(appIDOrName string, nodeNameOrID string) (*Node, error) {

	url := fmt.Sprintf("controller/rest/applications/%s/nodes/%s?output=json", appIDOrName, nodeNameOrID)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var node *Node
	err = s.client.Do(req, node)
	if err != nil {
		return nil, err
	}

	return node, nil
}
