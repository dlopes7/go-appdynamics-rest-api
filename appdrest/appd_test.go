package appdrest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestNewClient(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	_, err := appdrest.NewClient("invalid_protocol", os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	if err != nil {
		t.Logf("Expected error creating client with invalid port: %s", err.Error())
	}
}
