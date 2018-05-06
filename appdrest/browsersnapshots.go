package appdrest

// BrowserSnapshot represents a browser snapshot
type BrowserSnapshot struct {
	GeoCity                          string      `json:"geoCity"`
	GeoCountry                       string      `json:"geoCountry"`
	GeoRegion                        string      `json:"geoRegion"`
	IP                               string      `json:"ip"`
	PageName                         string      `json:"pageName"`
	PageType                         string      `json:"pageType"`
	Browser                          string      `json:"browser"`
	BrowserVersion                   string      `json:"browserVersion"`
	Device                           string      `json:"device"`
	DeviceOS                         string      `json:"deviceOS"`
	BtData                           interface{} `json:"btData"`
	AjaxErrorCode                    interface{} `json:"ajaxErrorCode"`
	XhrStatus                        interface{} `json:"xhrStatus"`
	PageTitle                        interface{} `json:"pageTitle"`
	PageURL                          string      `json:"pageUrl"`
	PageReferrer                     interface{} `json:"pageReferrer"`
	ClientRequestGUID                string      `json:"clientRequestGuid"`
	OtherClientRequestGuids          interface{} `json:"otherClientRequestGuids"`
	ServerSnapshotCg                 interface{} `json:"serverSnapshotCg"`
	ServerSnapshotError              interface{} `json:"serverSnapshotError"`
	BasePageGUID                     string      `json:"basePageGuid"`
	ParentGUID                       string      `json:"parentGuid"`
	ParentPageName                   string      `json:"parentPageName"`
	ParentPageType                   string      `json:"parentPageType"`
	ParentPageURL                    string      `json:"parentPageUrl"`
	TruncationMessages               interface{} `json:"truncationMessages"`
	SynthMeasurementID               interface{} `json:"synthMeasurementId"`
	SynthScheduleID                  interface{} `json:"synthScheduleId"`
	VirtualPageDigestCycles          int         `json:"virtualPageDigestCycles"`
	Synthetic                        bool        `json:"synthetic"`
	EndUserTime                      int         `json:"endUserTime"`
	FirstByteTime                    int         `json:"firstByteTime"`
	FrontEndTime                     int         `json:"frontEndTime"`
	DocumentReadyTime                int         `json:"documentReadyTime"`
	PageRenderTime                   int         `json:"pageRenderTime"`
	ResponseTime                     int         `json:"responseTime"`
	PageReadyTime                    int         `json:"pageReadyTime"`
	UserExperience                   string      `json:"userExperience"`
	PageLoadError                    string      `json:"pageLoadError"`
	ErrorOccurred                    bool        `json:"errorOccurred"`
	SyntheticEventStatus             interface{} `json:"syntheticEventStatus"`
	ID                               int         `json:"id"`
	ApplicationID                    int         `json:"applicationId"`
	Time                             int64       `json:"time"`
	PageAddID                        int         `json:"pageAddId"`
	ParentPageAddID                  int         `json:"parentPageAddId"`
	ResponseTimeDisplayValue         int         `json:"responseTimeDisplayValue"`
	BaseSnapshot                     interface{} `json:"baseSnapshot"`
	ParentSnapshot                   interface{} `json:"parentSnapshot"`
	ApplicationServerTime            int         `json:"applicationServerTime"`
	IsArchived                       bool        `json:"isArchived"`
	JavaScriptErrorString            string      `json:"javaScriptErrorString"`
	BtEventData                      interface{} `json:"btEventData"`
	ErrorEventData                   interface{} `json:"errorEventData"`
	AjaxError                        interface{} `json:"ajaxError"`
	CustomUserData                   interface{} `json:"customUserData"`
	ResourceTimingSnapshotDescriptor interface{} `json:"resourceTimingSnapshotDescriptor"`
	SyntheticScheduleName            interface{} `json:"syntheticScheduleName"`
}

type browserSnapshotList struct {
	Snapshots []*BrowserSnapshot `json:"snapshots"`
}

type browserSnapshotPost struct {
	ApplicationID   int    `json:"applicationId"`
	TimeRangeString string `json:"timeRangeString"`
}

// BrowserSnapshotService intermediates BrowserSnapshot requests
type BrowserSnapshotService service

// GetBrowserSnapshots obtains all browser snapshots from an Application
func (s *BrowserSnapshotService) GetBrowserSnapshots(appID int) ([]*BrowserSnapshot, error) {

	url := "/controller/restui/browserSnapshotList/getSnapshots"

	browserSnapshotPost := &browserSnapshotPost{
		ApplicationID:   41,
		TimeRangeString: "last_15_minutes.BEFORE_NOW.-1.-1.1",
	}

	var browserSnapshotList *browserSnapshotList
	err := s.client.Rest("POST", url, &browserSnapshotList, browserSnapshotPost)
	if err != nil {
		return nil, err
	}

	snapshots := browserSnapshotList.Snapshots

	return snapshots, nil
}
