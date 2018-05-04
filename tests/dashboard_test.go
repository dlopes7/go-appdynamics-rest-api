package tests

import (
	"testing"
)

func TestGetDashboards(t *testing.T) {
	client := CreateClient()
	_, err := client.Dashboard.GetDashboards()
	if err != nil {
		t.Errorf("Error getting dasboards: %s\n", err.Error())
		t.FailNow()
	}

}
