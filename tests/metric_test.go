package tests

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetMetricData(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))

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

	}
}
