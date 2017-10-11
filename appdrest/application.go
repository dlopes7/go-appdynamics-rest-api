package appdrest

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

	url := "rest/applications?output=json"

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
