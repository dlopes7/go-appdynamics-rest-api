package examples

/*
This example will list every single metric for an Application
This works with a recursive call to the GetMetricHierarchy function
*/

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func recursivePath(client *appdrest.Client, appName string, prefix string) {
	metricH, err := client.MetricData.GetMetricHierarchy(appName, prefix)
	if err != nil {
		panic(err.Error())
	}
	if prefix != "" {
		prefix = prefix + "|"
	}

	for _, metric := range metricH {
		if metric.Type == "folder" {
			// If we are in a folder, go down one more level
			recursivePath(client, appName, prefix+metric.Name)
		} else {
			// If we are here we got to the bottom, we have a metric
			fmt.Println(prefix + metric.Name)
		}
	}

}

func printMetricHierarchy() {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))

	apps, err := client.Application.GetApplications()
	if err != nil {
		panic(err.Error())
	}

	app := apps[0]
	recursivePath(client, app.Name, "")
}
