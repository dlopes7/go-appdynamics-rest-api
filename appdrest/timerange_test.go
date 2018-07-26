package appdrest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/dlopes7/go-appdynamics-rest-api/appdrest"
)

func TestTimeRanges(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("APPD_CONTROLLER_PORT"))
	client, _ := appdrest.NewClient(os.Getenv("APPD_CONTROLLER_PROTOCOL"), os.Getenv("APPD_CONTROLLER_HOST"), port, os.Getenv("APPD_USER"), os.Getenv("APPD_PASSWORD"), os.Getenv("APPD_ACCOUNT"))

	// Test GetTimeRanges
	trs, err := client.TimeRange.GetTimeRanges()
	if err != nil {
		t.Errorf("Error getting timeranges: %s\n", err.Error())
		t.FailNow()
	}
	if len(trs) > 0 {

		// Test GetTimeRangeByName
		tr, err := client.TimeRange.GetTimeRangeByName(trs[0].Name)
		if err != nil {
			t.Errorf("Error getting timerange %s: %s\n", trs[0].Name, err.Error())
			t.FailNow()
		}

		// Test UpdateTimeRange
		_, err = client.TimeRange.UpdateTimeRange(*tr)

	}

}
