package tests

import (
	"testing"
)

func TestGetMyAccount(t *testing.T) {
	client := CreateClient()
	acc, err := client.Account.GetMyAccount()
	if err != nil {
		t.Errorf("Error getting account: %s\n", err.Error())
		t.FailNow()
	}
	if len(acc.Name) == 0 {
		t.Errorf("Error getting account name: %v\n", acc)
		t.FailNow()
	}
}
