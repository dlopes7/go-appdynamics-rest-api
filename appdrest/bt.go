package appdrest

import (
	"fmt"
)

// BusinessTransaction represents one BT within one Application
type BusinessTransaction struct {
	InternalName   string `json:"internalName"`
	TierID         int    `json:"tierId"`
	EntryPointType string `json:"entryPointType"`
	Background     bool   `json:"background"`
	TierName       string `json:"tierName"`
	Name           string `json:"name"`
	ID             int    `json:"id"`
}

// BusinessTransactionService intermediates BusinessTransaction requests
type BusinessTransactionService service

// GetBusinessTransactions obtains all BTs from an application
func (s *BusinessTransactionService) GetBusinessTransactions(appID int) ([]*BusinessTransaction, error) {

	url := fmt.Sprintf("rest/applications/%d/business-transactions?output=json", appID)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var bts []*BusinessTransaction
	err = s.client.Do(req, &bts)
	if err != nil {
		return nil, err
	}

	return bts, nil
}
