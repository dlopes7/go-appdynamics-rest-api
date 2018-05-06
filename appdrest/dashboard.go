package appdrest

import (
	"fmt"
)

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
	Widgets                   []Widget      `json:"widgets"`
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

// Widget - One Dashboard contains multiple Widgets
type Widget struct {
	Type                        string        `json:"type"`
	ID                          int           `json:"id"`
	Version                     int           `json:"version"`
	GUID                        string        `json:"guid"`
	Title                       string        `json:"title"`
	DashboardID                 int           `json:"dashboardId"`
	WidgetsMetricMatchCriterias interface{}   `json:"widgetsMetricMatchCriterias"`
	Height                      int           `json:"height"`
	Width                       int           `json:"width"`
	MinHeight                   int           `json:"minHeight"`
	MinWidth                    int           `json:"minWidth"`
	X                           int           `json:"x"`
	Y                           int           `json:"y"`
	Label                       string        `json:"label"`
	Description                 string        `json:"description"`
	DrillDownURL                string        `json:"drillDownUrl"`
	UseMetricBrowserAsDrillDown bool          `json:"useMetricBrowserAsDrillDown"`
	DrillDownActionType         interface{}   `json:"drillDownActionType"`
	BackgroundColor             int           `json:"backgroundColor"`
	Color                       int           `json:"color"`
	FontSize                    int           `json:"fontSize"`
	UseAutomaticFontSize        bool          `json:"useAutomaticFontSize"`
	BorderEnabled               bool          `json:"borderEnabled"`
	BorderThickness             int           `json:"borderThickness"`
	BorderColor                 int           `json:"borderColor"`
	BackgroundAlpha             float64       `json:"backgroundAlpha"`
	ShowValues                  bool          `json:"showValues"`
	BackgroundColors            interface{}   `json:"backgroundColors"`
	CompactMode                 bool          `json:"compactMode"`
	ShowTimeRange               bool          `json:"showTimeRange"`
	RenderIn3D                  bool          `json:"renderIn3D"`
	ShowLegend                  bool          `json:"showLegend"`
	LegendPosition              interface{}   `json:"legendPosition"`
	LegendColumnCount           interface{}   `json:"legendColumnCount"`
	StartTime                   interface{}   `json:"startTime"`
	EndTime                     interface{}   `json:"endTime"`
	MinutesBeforeAnchorTime     int           `json:"minutesBeforeAnchorTime"`
	IsGlobal                    bool          `json:"isGlobal"`
	Properties                  []interface{} `json:"properties"`
	MissingEntities             interface{}   `json:"missingEntities"`
	AdqlQueries                 []string      `json:"adqlQueries"`
	AnalyticsType               string        `json:"analyticsType"`
	SearchMode                  string        `json:"searchMode"`
	IsStackingEnabled           bool          `json:"isStackingEnabled"`
	LegendsLayout               string        `json:"legendsLayout"`
	MaxAllowedYAxisFields       int           `json:"maxAllowedYAxisFields"`
	MaxAllowedXAxisFields       int           `json:"maxAllowedXAxisFields"`
	Min                         interface{}   `json:"min"`
	Max                         interface{}   `json:"max"`
	MinType                     interface{}   `json:"minType"`
	MaxType                     interface{}   `json:"maxType"`
	ShowMinExtremes             bool          `json:"showMinExtremes"`
	ShowMaxExtremes             bool          `json:"showMaxExtremes"`
	IntervalType                interface{}   `json:"intervalType"`
	Interval                    interface{}   `json:"interval"`
	DisplayPercentileMarkers    bool          `json:"displayPercentileMarkers"`
	PercentileValue1            interface{}   `json:"percentileValue1"`
	PercentileValue2            interface{}   `json:"percentileValue2"`
	PercentileValue3            interface{}   `json:"percentileValue3"`
	PercentileValue4            interface{}   `json:"percentileValue4"`
	Resolution                  interface{}   `json:"resolution"`
	DataFetchSize               interface{}   `json:"dataFetchSize"`
	PercentileLine              interface{}   `json:"percentileLine"`
	TimeRangeInterval           interface{}   `json:"timeRangeInterval"`
	PollingInterval             interface{}   `json:"pollingInterval"`
	Unit                        int           `json:"unit"`
	IsRawQuery                  bool          `json:"isRawQuery"`
	ViewState                   interface{}   `json:"viewState"`
	GridState                   interface{}   `json:"gridState"`
}

// DashboardService intermediates Dashboard requests
type DashboardService service

// GetDashboards obtains all dashboards from a controller
func (s *DashboardService) GetDashboards() ([]*Dashboard, error) {

	url := "/controller/restui/dashboards/getAllDashboardsByType/false"
	var dashboards []*Dashboard
	err := s.client.RestInternal("GET", url, &dashboards, nil)
	if err != nil {
		return nil, err
	}

	return dashboards, nil
}

// GetDashboard obtains a single dashboard from a controller
func (s *DashboardService) GetDashboard(ID int) (*Dashboard, error) {

	url := fmt.Sprintf("/controller/restui/dashboards/dashboardIfUpdated/%d/-1", ID)

	var dashboard *Dashboard
	err := s.client.RestInternal("GET", url, &dashboard, nil)
	if err != nil {
		return nil, err
	}

	return dashboard, nil
}

// TODO: Implement UpdateWidget
// UpdateWidget updates a Widget of a Dashboard ID
// func (s *DashboardService) UpdateWidget(dashboardID int, widget *Widget) {
// /controller/restui/dashboards/updateWidget
// }
