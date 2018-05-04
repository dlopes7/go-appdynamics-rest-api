package tests

import (
	"testing"
	"time"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetSnapshots(t *testing.T) {
	client := CreateClient()
	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting apps: %s\n", err.Error())
		t.FailNow()
	}
	if len(apps) > 0 {
		app := apps[0]
		filters := &appdrest.SnapshotFilters{}
		snapshots, err := client.Snapshot.GetSnapshots(app.ID, appdrest.TimeBEFORENOW, 15, time.Now(), time.Now(), filters)
		if err != nil {
			t.Errorf("Error getting snapshots: %s\n", err.Error())
			t.FailNow()
		}
		t.Logf("Got %d snapshots from app %s", len(snapshots), app.Name)

	}

}
