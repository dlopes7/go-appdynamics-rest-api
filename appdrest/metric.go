package appdrest

import (
	"fmt"
	"time"
)

// MetricData contains metric values for a single metric
type MetricData struct {
	MetricName   string        `json:"metricName"`
	MetricID     int           `json:"metricId"`
	MetricPath   string        `json:"metricPath"`
	Frequency    string        `json:"frequency"`
	MetricValues []MetricValue `json:"metricValues"`
}

// MetricValue is always part of an array of metrics, inside a MetricData struct
type MetricValue struct {
	Occurrences       int   `json:"occurrences"`
	Current           int   `json:"current"`
	Min               int   `json:"min"`
	Max               int   `json:"max"`
	StartTimeInMillis int64 `json:"startTimeInMillis"`
	UseRange          bool  `json:"useRange"`
	Count             int   `json:"count"`
	Sum               int   `json:"sum"`
	Value             int   `json:"value"`
	StandardDeviation int   `json:"standardDeviation"`
}

// Metric represents a Metric object that might be a folder or child
type Metric struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Consts for the technique used to obtain metric data
const (
	TimeBEFORENOW    = "BEFORE_NOW"
	TimeBEFORETIME   = "BEFORE_TIME"
	TimeAFTERTIME    = "AFTER_TIME"
	TimeBETWEENTIMES = "BETWEEN_TIMES"
)

// MetricDataService intermediates MetricData requests
type MetricDataService service

// GetMetricData obtains metrics matching a pattern
func (s *MetricDataService) GetMetricData(appID int, metricPath string, rollup bool, timeRangeType string, durationInMins int, startTime time.Time, endTime time.Time) ([]*MetricData, error) {

	url := fmt.Sprintf("controller/rest/applications/%d/metric-data?output=json", appID)
	url += fmt.Sprintf("&rollup=%t", rollup)

	url += fmt.Sprintf("&metric-path=%s", metricPath)
	url += fmt.Sprintf("&time-range-type=%s", timeRangeType)

	if timeRangeType == TimeBEFORENOW || timeRangeType == TimeBEFORETIME || timeRangeType == TimeAFTERTIME {
		url += fmt.Sprintf("&duration-in-mins=%d", durationInMins)

	}
	if timeRangeType == TimeAFTERTIME || timeRangeType == TimeBETWEENTIMES {
		url += fmt.Sprintf("&start-time=%v", startTime)
	}
	if timeRangeType == TimeBEFORETIME || timeRangeType == TimeBETWEENTIMES {
		url += fmt.Sprintf("&end-time=%v", endTime)
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var metrics []*MetricData
	err = s.client.Do(req, &metrics)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

// GetMetricHierarchy obtains the Metric Browser hierarchy
func (s *MetricDataService) GetMetricHierarchy(appID int, metricPath string) ([]*Metric, error) {
	url := fmt.Sprintf("controller/rest/applications/%d/metrics?output=json", appID)

	if metricPath != "" {
		url += fmt.Sprintf("&metric-path=%s", metricPath)
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var metrics []*Metric
	err = s.client.Do(req, &metrics)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}
