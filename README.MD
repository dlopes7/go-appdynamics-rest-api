# appdrest #
  
[![CircleCI](https://circleci.com/gh/dlopes7/go-appdynamics-rest-api.svg?style=svg)](https://circleci.com/gh/dlopes7/go-appdynamics-rest-api)  [![codecov](https://codecov.io/gh/dlopes7/go-appdynamics-rest-api/branch/master/graph/badge.svg)](https://codecov.io/gh/dlopes7/go-appdynamics-rest-api)

  
Library to access the [AppDynamics REST API](https://docs.appdynamics.com/display/PRO43/AppDynamics+APIs)  

## Usage ##
```go
import "github.com/dlopes7/go-appdynamics-rest-api/appdrest"
```

Create a client, get every Business Transaction for every Application

```go
client, _ :=  appdrest.NewClient("http", "192.168.33.10", 8090, "admin", "password", "customer1")

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
```



