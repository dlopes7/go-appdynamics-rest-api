package tests

import (
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"

	"os"
)

func TestGetApplications(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting applications: %s\n", err.Error())
		t.FailNow()
	}
	if len(apps) == 0 {
		t.Error("No applications were returned")
	}
}

func TestGetApplication(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting applications: %s\n", err.Error())
		t.FailNow()
	}
	if len(apps) == 0 {
		t.Error("No applications were returned")
	}

	app, err := client.Application.GetApplication(strconv.Itoa(apps[0].ID))
	if err != nil {
		t.Errorf("Error getting Application %s", apps[0].Name)
		t.FailNow()
	}
	if len(app.Name) == 0 || app.ID <= 0 {
		t.Errorf("Error getting Application properties for %s", apps[0].Name)
	}

	app, err = client.Application.GetApplication(apps[0].Name)
	if err != nil {
		t.Errorf("Error getting Application %s", apps[0].Name)
		t.FailNow()
	}
	if len(app.Name) == 0 || app.ID <= 0 {
		t.Errorf("Error getting Application properties for %s", apps[0].Name)
	}
}
