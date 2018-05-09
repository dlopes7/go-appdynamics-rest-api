package examples

import (
	"fmt"
	"time"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func getErrorSnapshots() {
	client, _ := appdrest.NewClient("http", "192.168.33.10", 8090, "admin", "password", "customer1")

	apps, err := client.Application.GetApplications()
	if err != nil {
		panic(err.Error())
	}

	snapsFilters := &appdrest.SnapshotFilters{
		ErrorOccurred: true,
	}

	snaps, err := client.Snapshot.GetSnapshots(apps[0].ID, appdrest.TimeBEFORENOW, 10, time.Now(), time.Now(), snapsFilters)

	for _, snap := range snaps {
		fmt.Println(snap.BusinessTransactionID, snap.Summary)
	}

}
