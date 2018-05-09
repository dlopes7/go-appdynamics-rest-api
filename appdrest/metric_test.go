package appdrest_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetMetricData(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))

	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting apps: %s\n", err.Error())
		t.FailNow()
	}
	if len(apps) > 0 {
		app := apps[0]
		metrics, err := client.MetricData.GetMetricData(app.Name, "Overall Application Performance|Calls per Minute", false, appdrest.TimeBEFORENOW, 60, time.Now(), time.Now())
		if err != nil {
			t.Errorf("Error getting metrics: %s\n", err.Error())
			t.FailNow()
		}
		if len(metrics) > 0 {
			t.Logf("Got %+v datapoints from app %s", len(metrics[0].MetricValues), app.Name)
		}

		// Test for TimeBETWEENTIMES
		past := time.Now().Add(time.Duration(-15) * time.Minute)
		metrics, err = client.MetricData.GetMetricData(app.Name, "Overall Application Performance|Calls per Minute", false, appdrest.TimeBETWEENTIMES, 0, past, time.Now())
		if err != nil {
			t.Errorf("Error getting metrics: %s\n", err.Error())
			t.FailNow()
		}
		if len(metrics) > 0 {
			t.Logf("Got %+v datapoints from app %s", len(metrics[0].MetricValues), app.Name)
		}

		// Test for Invalid Metric
		_, err = client.MetricData.GetMetricData("Invalid Application Name", "Overall Application Performance|Calls per Minute", false, appdrest.TimeBETWEENTIMES, 60, time.Now(), time.Now())
		if err != nil {
			t.Logf("Expected error getting metrics for invalid Application: %s\n", err.Error())
		}

	}
}

func TestGetMetricHierarchy(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))

	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting apps: %s\n", err.Error())
		t.FailNow()
	}

	app := apps[0]
	metricH, err := client.MetricData.GetMetricHierarchy(app.Name, "")
	if err != nil {
		t.Errorf("Error getting metric hierarchy: %s\n", err.Error())
		t.FailNow()
	}

	t.Logf("Found %d metric folders for App %s", len(metricH), app.Name)

	metricH, err = client.MetricData.GetMetricHierarchy(app.Name, metricH[0].Name)
	if err != nil {
		t.Errorf("Error getting metric hierarchy: %s\n", err.Error())
		t.FailNow()
	}

	_, err = client.MetricData.GetMetricHierarchy("Invalid App Name", "")
	if err != nil {
		t.Logf("Expected error getting metric hierarchy for invalid Application: %s\n", err.Error())
	}
}
