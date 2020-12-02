package tests

import (
	"secureauth/internal/logs"
	"testing"
)

func TestGetRealms(t *testing.T) {

	// Store sample realm data
	confApp1 := "Salesforce"
	confFile1 := "resources/Audit1.Log"
	confTime1 := "2000-01-01T11:11:11.000Z"
	confRealm1 := "SecureAuth999"

	confApp2 := "AWS"
	confFile2 := "resources/Audit2.Log"
	confTime2 := "2000-02-02T22:22:22.000Z"
	confRealm2 := "SecureAuth1000"

	// Retrieve realms configurations
	realms := logs.GetRealms("./resources")

	// Validate the returned configuration count
	wantedConfigCount := 2
	retrievedConfigCount := len(realms)
	if retrievedConfigCount != wantedConfigCount {
		t.Errorf("The number of retrieved configurations did not match the expected value. Retrieved: %d. Wanted: %d.", retrievedConfigCount, wantedConfigCount)
	}

	// Validate the first realms values
	retrievedApp1 := realms[0].GetString("application")
	retrievedFile1 := realms[0].GetString("file")
	retrievedTime1 := realms[0].GetString("latest_time")
	retrievedRealm1 := realms[0].GetString("realm")

	if retrievedApp1 != confApp1 {
		t.Errorf("The application value retrieved from the config file does not match the expected value. Retrieved: %s. Wanted: %s.", retrievedApp1, confApp1)
	}

	if retrievedFile1 != confFile1 {
		t.Errorf("The file value retrieved from the config file does not match the expected value. Retrieved: %s. Wanted: %s.", retrievedFile1, confFile1)
	}

	if retrievedTime1 != confTime1 {
		t.Errorf("The time value retrieved from the config file does not match the expected value. Retrieved: %s. Wanted: %s.", retrievedTime1, confTime1)
	}

	if retrievedRealm1 != confRealm1 {
		t.Errorf("The realm value retrieved from the config file does not match the expected value. Retrieved: %s. Wanted: %s.", retrievedRealm1, confRealm1)
	}

	// Validate the second realms values
	retrievedApp2 := realms[1].GetString("application")
	retrievedFile2 := realms[1].GetString("file")
	retrievedTime2 := realms[1].GetString("latest_time")
	retrievedRealm2 := realms[1].GetString("realm")

	if retrievedApp2 != confApp2 {
		t.Errorf("The application value retrieved from the config file does not match the expected value. Retrieved: %s. Wanted: %s.", retrievedApp2, confApp2)
	}

	if retrievedFile2 != confFile2 {
		t.Errorf("The file value retrieved from the config file does not match the expected value. Retrieved: %s. Wanted: %s.", retrievedFile2, confFile2)
	}

	if retrievedTime2 != confTime2 {
		t.Errorf("The time value retrieved from the config file does not match the expected value. Retrieved: %s. Wanted: %s.", retrievedTime2, confTime2)
	}

	if retrievedRealm2 != confRealm2 {
		t.Errorf("The realm value retrieved from the config file does not match the expected value. Retrieved: %s. Wanted: %s.", retrievedRealm2, confRealm2)
	}

}