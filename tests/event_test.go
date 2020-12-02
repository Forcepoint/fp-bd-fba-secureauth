package tests

import (
	"fmt"
	"secureauth/internal/structs"
	"testing"
)

func TestCreateEvent(t *testing.T) {

	// Sample Values
	sampleDetails := [][]string{
		{"LogChannel=\"Test-Channel\"", "LogChannel", "Test-Channel"},
		{"FormatVersion=\"1.0.0\"", "FormatVersion", "1.0.0"},
		{"EventID=\"51170\"", "EventID", "51170"},
		{"Timestamp=\"1995-07-17T15:40:00.000Z\"", "Timestamp", "1995-07-17T15:40:00.000Z"},
		{"CompanyID=\"\"", "CompanyID"},
		{"ApplianceID=\"\"", "ApplianceID"},
		{"Realm=\"test-realm\"", "Realm", "test-realm"},
		{"UserID=\"John.McClane\"", "UserID", "John.McClane"},
		{"BrowserSession=\"128d76a2-b5b1-4e36-9a2f-a7a3a1890ad3\"", "BrowserSession", "128d76a2-b5b1-4e36-9a2f-a7a3a1890ad3"},
		{"RequestID=\"68aa2261-9adb-4b21-9b12-533696ab4888\"", "RequestID", "68aa2261-9adb-4b21-9b12-533696ab4888"},
		{"UserHostAddress=\"234.123.34.56\"", "UserHostAddress", "234.123.34.56"},
	}
	sampleEventTypes := map[string]string {
		"51170": "Password validated. Successful login.",
	}
	sampleApp := "test-app"
	sampleRealm := "test-realm"

	// Create event using test details
	event := structs.CreateEvent(sampleDetails, sampleApp, sampleRealm)

	// Validate the event type values are correct
	wantedType := "Login Attempt"
	retrievedType := event.Type
	if retrievedType != wantedType {
		t.Errorf("The value for Event Type was incorrect. Retrieved: %s. Wanted: %s.", retrievedType, wantedType)
	}

	// Validate that the timestamp is correct
	wantedTimestamp := sampleDetails[3][2]
	retrievedTimestamp := event.Timestamp
	if retrievedTimestamp != wantedTimestamp {
		t.Errorf("The value for Event Timestamp was incorrect. Retrieved: %s. Wanted: %s.", retrievedTimestamp, wantedTimestamp)
	}

	// Validate length of event entities
	wantedEntitiesLength := 5
	retrievedEntitiesLength := len(event.Entities)
	if retrievedEntitiesLength != wantedEntitiesLength {
		t.Errorf("The length of the event entities array did not match. Retrieved: %d. Wanted: %d.", retrievedEntitiesLength, wantedEntitiesLength)
	}

	// Validate the event entities values
	wantedVendor := "SecureAuth"
	retrievedVendor := event.Entities[0].Entities[0]
	if retrievedVendor != wantedVendor {
		t.Errorf("The value for the Vendor entity was incorrect. Retrieved: %s. Wanted: %s.", retrievedVendor, wantedVendor)
	}

	wantedApplication := sampleApp
	retrievedApplication := event.Entities[1].Entities[0]
	if retrievedApplication != wantedApplication {
		t.Errorf("The value for the Application entity was incorrect. Retrieved: %s. Wanted: %s.", retrievedApplication, wantedApplication)
	}

	wantedRealm := sampleRealm
	retrievedRealm := event.Entities[2].Entities[0]
	if retrievedRealm != wantedRealm {
		t.Errorf("The value for the " +
			"Realm entity was incorrect. Retrieved: %s. Wanted: %s.", retrievedRealm, wantedRealm)
	}

	wantedSourceIP := sampleDetails[10][2]
	retrievedSourceIP := event.Entities[3].Entities[0]
	if retrievedSourceIP != wantedSourceIP {
		t.Errorf("The value for the Source IP entity was incorrect. Retrieved: %s. Wanted: %s.", retrievedSourceIP, wantedSourceIP)
	}

	wantedUser := sampleDetails[7][2]
	retrievedUser := event.Entities[4].Entities[0]
	if retrievedUser != wantedUser {
		t.Errorf("The value for the User entity was incorrect. Retrieved: %s. Wanted: %s.", retrievedUser, wantedUser)
	}

	// Validate the source event ID value
	wantedEventID := fmt.Sprintf("%s%s", sampleDetails[2][2], sampleDetails[9][2])
	retrievedEventID := event.SourceEventID
	if retrievedEventID != wantedEventID {
		t.Errorf("The value for the Source Event ID was incorrect. Retrieved: %s. Wanted: %s.", retrievedEventID, wantedEventID)
	}

	// Validate the subject value
	wantedSubject := sampleEventTypes[sampleDetails[2][2]]
	retrievedSubject := event.Subject
	if retrievedSubject != wantedSubject {
		t.Errorf("The value for the Subject was incorrect. Retrieved: %s. Wanted: %s.", retrievedSubject, wantedSubject)
	}

	// Validate the label value
	wantedLabel := "policy"
	retrievedLabel := event.Labels[0]
	if retrievedLabel != wantedLabel {
		t.Errorf("The value for label was incorrect. Retrieved: %s. Wanted: %s", retrievedLabel, wantedLabel)
	}

	// Validate the length of the attributes array
	wantedAttributesLength := 4
	retrievedAttributesLength := len(event.Attributes)
	if retrievedAttributesLength != wantedAttributesLength {
		t.Errorf("The value for attributes array length was incorrect. Retrieved: %d. Wanted: %d.", retrievedEntitiesLength, wantedAttributesLength)
	}

}
