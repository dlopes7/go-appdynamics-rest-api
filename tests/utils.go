package tests

import (
	"os"
	"strconv"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

// CreateClient creates a test client using environment variables
func CreateClient() *appdrest.Client {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	return client
}
