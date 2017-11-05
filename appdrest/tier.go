package appdrest

import (
	"fmt"
)

// Tier represents one tier within one Application
type Tier struct {
	AgentType     string `json:"agentType"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ID            int    `json:"id"`
	NumberOfNodes int    `json:"numberOfNodes"`
	Type          string `json:"type"`
}

// TierService intermediates Tier requests
type TierService service

// GetTiers obtains all Tiers from an Application
func (s *TierService) GetTiers(appID int) ([]*Tier, error) {

	url := fmt.Sprintf("controller/rest/applications/%d/tiers?output=json", appID)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var tiers []*Tier
	err = s.client.Do(req, &tiers)
	if err != nil {
		return nil, err
	}

	return tiers, nil
}
