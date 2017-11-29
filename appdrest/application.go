package appdrest

import (
	"fmt"
)

// Application represents a single Business Application within AppDynamics
type Application struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          int    `json:"id"`
}

// ApplicationService intermediates Application requests
type ApplicationService service

// GetApplications obtains all applications from a controller
func (s *ApplicationService) GetApplications() ([]*Application, error) {

	url := "controller/rest/applications?output=json"

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var apps []*Application
	err = s.client.Do(req, &apps)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

// GetApplication gets an Application by Name or ID
func (s *ApplicationService) GetApplication(appNameOrID string) (*Application, error) {

	url := fmt.Sprintf("controller/rest/applications/%s?output=json", appNameOrID)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var apps []*Application
	err = s.client.Do(req, &apps)
	if err != nil {
		return nil, err
	}

	return apps[0], nil
}
