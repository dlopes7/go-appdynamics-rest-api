package appdrest_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetSnapshots(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
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

func TestGetSnapshotsFiltered(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting apps: %s\n", err.Error())
		t.FailNow()
	}
	if len(apps) > 0 {
		app := apps[0]
		filters := &appdrest.SnapshotFilters{
			ExecutionTimeInMilis: 3000,
			MaximumResults:       10,
			UserExperience:       []string{"SLOW", "ERROR", "VERY_SLOW"},
		}
		then := time.Now().Add(time.Duration(-20) * time.Minute)
		thenMinus5 := time.Now().Add(time.Duration(-5) * time.Minute)

		snapshots, err := client.Snapshot.GetSnapshots(app.ID, appdrest.TimeAFTERTIME, 15, then, thenMinus5, filters)
		if err != nil {
			t.Errorf("Error getting snapshots: %s\n", err.Error())
			t.FailNow()
		}
		t.Logf("Got %d snapshots from app %s", len(snapshots), app.Name)

	}

}
