package appdrest_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetMyAccount(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	acc, err := client.Account.GetMyAccount()
	if err != nil {
		t.Errorf("Error getting account: %s\n", err.Error())
		t.FailNow()
	}
	t.Logf("Got account %s", acc.Name)

	// Test with an invalid client
	client, _ = appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), 1, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	_, err = client.Account.GetMyAccount()
	if err != nil {
		t.Logf("Expected error getting account with invalid client: %s\n", err.Error())
	}
}

func TestGetLicenseModules(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	acc, err := client.Account.GetMyAccount()
	if err != nil {
		t.Errorf("Error getting account: %s\n", err.Error())
		t.FailNow()
	}
	if len(acc.Name) == 0 {
		t.Errorf("Error getting account name: %v\n", acc)
		t.FailNow()
	}
	modules, err := client.Account.GetLicenseModules(acc.ID)
	if err != nil {
		t.Errorf("Error getting license Modules: %s\n", err.Error())
		t.FailNow()
	}
	t.Logf("Got properties for %d license modules", len(modules))

	_, err = client.Account.GetLicenseModules("-1")
	if err != nil {
		t.Logf("Expeceted error getting license Modules for invalid Account: %s\n", err.Error())
	}

}

func TestGetLicenseProperties(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))

	acc, err := client.Account.GetMyAccount()
	if err != nil {
		t.Errorf("Error getting account: %s\n", err.Error())
		t.FailNow()
	}
	if len(acc.Name) == 0 {
		t.Errorf("Error getting account name: %v\n", acc)
		t.FailNow()
	}

	properties, err := client.Account.GetLicenseProperties(acc.ID, "apm")
	if err != nil {
		t.Errorf("Error getting license Properties: %s\n", err.Error())
		t.FailNow()
	}

	t.Logf("Agent type APM, %s: %s", properties[2].Name, properties[2].Value)

	_, err = client.Account.GetLicenseProperties(acc.ID, "doesnt_exist")
	if err != nil {
		t.Logf("Expected error when getting invalid license Properties: %s\n", err.Error())
	}
}

func TestGetLicenseUsages(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))

	acc, err := client.Account.GetMyAccount()
	if err != nil {
		t.Errorf("Error getting account: %s\n", err.Error())
		t.FailNow()
	}
	if len(acc.Name) == 0 {
		t.Errorf("Error getting account name: %v\n", acc)
		t.FailNow()
	}

	usages, err := client.Account.GetLicenseUsages(acc.ID, "apm")
	if err != nil {
		if strings.Contains(err.Error(), "500") {
			t.Logf("User doesn't have permission to check license usage")
		} else {
			t.Errorf("Error getting license Usages: %s\n", err.Error())
		}
	}
	for usage := range usages {
		t.Logf("Usage %+v", usage)
	}

}
