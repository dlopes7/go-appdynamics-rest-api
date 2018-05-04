package tests

import (
	"testing"
)

func TestGetTiers(t *testing.T) {
	client := CreateClient()
	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting apps: %s\n", err.Error())
		t.FailNow()
	}
	if len(apps) > 0 {
		app := apps[0]
		tiers, err := client.Tier.GetTiers(app.ID)
		if err != nil {
			t.Errorf("Error getting applications: %s\n", err.Error())
			t.FailNow()
		}
		t.Logf("Got %d tiers from app %s", len(tiers), app.Name)
	}
}
