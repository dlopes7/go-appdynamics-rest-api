package appdrest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetDashboards(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	_, err := client.Dashboard.GetDashboards()
	if err != nil {
		t.Errorf("Error getting dasboards: %s\n", err.Error())
		t.FailNow()
	}

	// Test invalid client
	client, _ = appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), 1, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	_, err = client.Dashboard.GetDashboards()
	if err != nil {
		t.Logf("Expected error using invalid client: %s\n", err.Error())
	}

}

func TestGetDashboard(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	dashboards, err := client.Dashboard.GetDashboards()
	if err != nil {
		t.Errorf("Error getting dasboards: %s\n", err.Error())
		t.FailNow()
	}
	if len(dashboards) == 0 {
		t.Log("Found no dashboards, not testing GetDashboard")
	} else {
		dashboard_id := dashboards[0].ID
		dashboard, err := client.Dashboard.GetDashboard(dashboard_id)
		if err != nil {
			t.Errorf("Error getting dasboard ID %d: %s\n", dashboard_id, err.Error())
			t.FailNow()
		}
		t.Logf("Found dashboard '%s' with %d widgets", dashboard.Name, len(dashboard.Widgets))
	}

	// Test for invalid dashboard ID
	_, err = client.Dashboard.GetDashboard(-1)
	if err != nil {
		t.Logf("Expected error for invalid dashboard ID. %s\n", err.Error())
	}

}
