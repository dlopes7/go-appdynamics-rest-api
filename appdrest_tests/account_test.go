package appdrest_tests

import (
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"

	"os"
)

func TestGetMyAccount(t *testing.T) {
	client := appdrest.NewClient("http", "demo2.appdynamics.com", 80, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
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
