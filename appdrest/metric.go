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

const (
	BeforeNow    = "BEFORE_NOW"
	BeforeTime   = "BEFORE_TIME"
	AfterTime    = "AFTER_TIME"
	BetweenTimes = "BETWEEN_TIMES"
)

// MetricDataService intermediates MetricData requests
type MetricDataService service

// GetMetricData obtains metrics matching a pattern
func (s *MetricDataService) GetMetricData(appID int, metricPath string, rollup bool, timeRangeType string, durationInMins int, startTime time.Time, endTime time.Time) ([]*MetricData, error) {

	url := fmt.Sprintf("rest/applications/%d/metric-data?output=json", appID)
	url += fmt.Sprintf("&rollup=%t", rollup)

	url += fmt.Sprintf("&metric-path=%s", metricPath)
	url += fmt.Sprintf("&time-range-type=%s", timeRangeType)

	if timeRangeType == BeforeNow || timeRangeType == BeforeTime || timeRangeType == AfterTime {
		url += fmt.Sprintf("&duration-in-mins=%d", durationInMins)

	}
	if timeRangeType == AfterTime || timeRangeType == BetweenTimes {
		url += fmt.Sprintf("&start-time=%v", startTime)
	}
	if timeRangeType == BeforeTime || timeRangeType == BetweenTimes {
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
