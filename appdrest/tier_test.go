package appdrest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetTiers(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting apps: %s\n", err.Error())
		t.FailNow()
	}
	if len(apps) > 0 {
		app := apps[0]
		tiers, err := client.Tier.GetTiers(app.ID)
		if err != nil {
			t.Errorf("Error getting tiers: %s\n", err.Error())
			t.FailNow()
		}
		t.Logf("Got %d tiers from app %s", len(tiers), app.Name)

		_, err = client.Tier.GetTiers(-1)
		if err != nil {
			t.Logf("Expected error getting tiers for invalid app: %s\n", err.Error())
		}
	}
}
