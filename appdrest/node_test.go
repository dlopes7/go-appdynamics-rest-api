package appdrest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestGetNodes(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))
	apps, err := client.Application.GetApplications()
	if err != nil {
		t.Errorf("Error getting apps: %s\n", err.Error())
		t.FailNow()
	}
	app := apps[0]

	nodes, err := client.Node.GetNodes(strconv.Itoa(app.ID))
	if err != nil {
		t.Errorf("Error getting Nodes: %s\n", err.Error())
		t.FailNow()
	}
	t.Logf("Got %d Nodes for App %s", len(nodes), app.Name)

	_, err = client.Node.GetNodes("-1")
	if err != nil {
		t.Logf("Expected error, getting Nodes for invalid Application: %s\n", err.Error())
	}

	node, err := client.Node.GetNode(strconv.Itoa(app.ID), strconv.Itoa(nodes[0].ID))
	if err != nil {
		t.Errorf("Error getting Node: %s\n", err.Error())
		t.FailNow()
	}
	t.Logf("Got Node %s for App %s", node.Name, app.Name)

	_, err = client.Node.GetNode(strconv.Itoa(app.ID), "-1")
	if err != nil {
		t.Logf("Expected error, getting Node with invalid ID: %s\n", err.Error())

	}

}
