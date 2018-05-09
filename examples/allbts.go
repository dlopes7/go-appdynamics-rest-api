package examples

import (
	"fmt"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func getAllBTs() {
	client, _ := appdrest.NewClient("http", "192.168.33.10", 8090, "admin", "password", "customer1")

	apps, err := client.Application.GetApplications()
	if err != nil {
		panic(err.Error())
	}
	for _, app := range apps {

		bts, err := client.BusinessTransaction.GetBusinessTransactions(app.ID)
		if err != nil {
			panic(err.Error())
		}
		for _, bt := range bts {
			fmt.Printf("App: %s, Tier: %s, BT: %s, Type: %s\n", app.Name, bt.TierName, bt.Name, bt.EntryPointType)
		}

	}

}
