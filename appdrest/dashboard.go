package appdrest

// Dashboard represents a single Dashboard within AppDynamics
type Dashboard struct {
	ID                        int           `json:"id"`
	Version                   int           `json:"version"`
	Name                      string        `json:"name"`
	NameUnique                bool          `json:"nameUnique"`
	BuiltIn                   bool          `json:"builtIn"`
	CreatedBy                 string        `json:"createdBy"`
	CreatedOn                 int64         `json:"createdOn"`
	ModifiedBy                string        `json:"modifiedBy"`
	ModifiedOn                int64         `json:"modifiedOn"`
	Description               interface{}   `json:"description"`
	MissingAssociatedEntities interface{}   `json:"missingAssociatedEntities"`
	Widgets                   interface{}   `json:"widgets"`
	SecurityToken             interface{}   `json:"securityToken"`
	SharingRevoked            bool          `json:"sharingRevoked"`
	WarRoom                   bool          `json:"warRoom"`
	Template                  bool          `json:"template"`
	TemplateEntityType        string        `json:"templateEntityType"`
	MinutesBeforeAnchorTime   int           `json:"minutesBeforeAnchorTime"`
	StartTime                 int           `json:"startTime"`
	EndTime                   int           `json:"endTime"`
	RefreshInterval           int           `json:"refreshInterval"`
	BackgroundColor           int           `json:"backgroundColor"`
	Color                     int           `json:"color"`
	Height                    int           `json:"height"`
	Width                     int           `json:"width"`
	Disabled                  bool          `json:"disabled"`
	CanvasType                string        `json:"canvasType"`
	LayoutType                string        `json:"layoutType"`
	Properties                []interface{} `json:"properties"`
}

// DashboardService intermediates Dashboard requests
type DashboardService service

// GetDashboards obtains all dashboards from a controller
func (s *DashboardService) GetDashboards() ([]*Dashboard, error) {

	url := "/controller/restui/dashboards/getAllDashboardsByType/false"

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var dashboards []*Dashboard
	err = s.client.DoRestUI(req, &dashboards)
	if err != nil {
		return nil, err
	}

	return dashboards, nil
}
