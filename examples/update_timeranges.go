package examples

import (
	"fmt"
	"time"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
	"github.com/jinzhu/now"
	"github.com/robfig/cron"
)

var client, err = appdrest.NewClient("https", "customer.saas.appdynamics.com", 443, "user", "password", "customer")

func updateTimeRange(name string, startTime time.Time, endTime time.Time) {
	timeRange, err := client.TimeRange.GetTimeRangeByName(name)

	if err != nil {
		fmt.Printf("Could not find the timerange: %s: %s\n", name, err.Error())
	}

	timeRange.ModifiedOn = time.Now().UnixNano() / 1000000
	timeRange.TimeRange.StartTime = startTime.UnixNano() / 1000000
	timeRange.TimeRange.EndTime = endTime.UnixNano() / 1000000

	_, err = client.TimeRange.UpdateTimeRange(*timeRange)
	if err != nil {
		fmt.Printf("Could not update the timerange %s: %s\n", name, err.Error())
	}

	fmt.Printf("Time Range '%s' updated\n", name)

}

func main() {

	c := cron.New()

	c.AddFunc("@every 1m", func() { updateTimeRange("Current Month", now.BeginningOfMonth(), time.Now()) })
	c.AddFunc("@every 1m", func() { updateTimeRange("Current Week", now.BeginningOfWeek(), time.Now()) })
	c.AddFunc("@every 30m", func() {
		updateTimeRange("Last Month", now.BeginningOfMonth().AddDate(0, -1, 0), now.BeginningOfMonth())
	})
	c.AddFunc("@every 30m", func() {
		updateTimeRange("Last Week", now.BeginningOfWeek().AddDate(0, 0, -7), now.BeginningOfWeek())
	})

	c.Start()

	select {}
}
