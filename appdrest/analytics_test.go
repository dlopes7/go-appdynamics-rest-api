package appdrest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetAnalyticsSearches(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))

	searches, err := client.Analytics.GetAnalyticsSearches()
	if err != nil {
		t.Errorf("Error getting analytics searches: %s\n", err.Error())
		t.FailNow()
	} else {
		t.Logf("Found %d searches", len(searches))
	}

	client, _ = appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), 1, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	_, err = client.Analytics.GetAnalyticsSearches()
	if err != nil {
		t.Logf("Expected error getting analytics searches for invalid client: %s\n", err.Error())
	}

}
