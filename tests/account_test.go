package tests

import (
	"os"
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetMyAccount(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	acc, err := client.Account.GetMyAccount()
	if err != nil {
		t.Errorf("Error getting account: %s\n", err.Error())
		t.FailNow()
	}
	if len(acc.Name) == 0 {
		t.Errorf("Error getting account name: %v\n", acc)
		t.FailNow()
	}
}
