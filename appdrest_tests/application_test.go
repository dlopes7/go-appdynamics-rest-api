package appdrest_tests

import (
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"

	"os"
)

var client = appdrest.NewClient("http", "demo2.appdynamics.com", 80, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))

func TestGetApplications(t *testing.T) {
	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting applications: %s\n", err.Error())
	}
	if len(apps) == 0 {
		t.Error("No applications were returned")
	}
}
