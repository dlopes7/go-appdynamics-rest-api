package appdrest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetBusinessTransactions(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting apps: %s\n", err.Error())
		t.FailNow()
	}
	app := apps[0]

	bts, err := client.BusinessTransaction.GetBusinessTransactions(app.ID)
	if err != nil {
		t.Errorf("Error getting BTs: %s\n", err.Error())
		t.FailNow()
	}
	t.Logf("Got %d BTs for App %s", len(bts), app.Name)

	_, err = client.BusinessTransaction.GetBusinessTransactions(-1)
	if err != nil {
		t.Logf("Expected error, getting BTs for invalid Application: %s\n", err.Error())
	}

}
