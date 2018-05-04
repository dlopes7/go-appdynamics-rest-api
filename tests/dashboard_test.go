package tests

import (
	"os"
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetDashboards(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	_, err := client.Dashboard.GetDashboards()
	if err != nil {
		t.Errorf("Error getting dasboards: %s\n", err.Error())
		t.FailNow()
	}

}
