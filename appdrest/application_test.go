package appdrest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetApplications(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting applications: %s\n", err.Error())
		t.FailNow()
	}
	if len(apps) == 0 {
		t.Error("No applications were returned")
	}

	client, _ = appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), 1, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	_, err = client.Application.GetApplications()
	if err != nil {
		t.Logf("Expected error getting applications for invalid client: %s\n", err.Error())
	}
}

func TestGetApplication(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
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

	// Invalid AppID
	app, err = client.Application.GetApplication("Invalid Application Name")
	if err != nil {
		t.Logf("Expected error getting invalid Application %s", err.Error())
	}
}

func TestGetApplicationsAllTypes(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	apps, err := client.Application.GetApplicationsAllTypes()
	if err != nil {
		t.Errorf("Error getting applications: %s\n", err.Error())
		t.FailNow()
	}
	if len(apps) == 0 {
		t.Error("No applications were returned")
	}

	client, _ = appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), 1, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	_, err = client.Application.GetApplicationsAllTypes()
	if err != nil {
		t.Logf("Expected error getting applications with invalid client: %s\n", err.Error())
	}
}

func TestExportApplicationConfig(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	apps, err := client.Application.GetApplicationsAllTypes()
	if err != nil {
		t.Errorf("Error getting applications: %s\n", err.Error())
		t.FailNow()
	}
	if len(apps) == 0 {
		t.Error("No applications were returned")
	}
	app := apps[0]
	_, err = client.Application.ExportApplicationConfig(app.ID)
	if err != nil {
		t.Errorf("Error exporting app %s: %s", app.Name, err.Error())
	}

	_, err = client.Application.ExportApplicationConfig(-1)
	if err != nil {
		t.Logf("Expected error trying to export invalid Application: %s", err.Error())
	}
}
