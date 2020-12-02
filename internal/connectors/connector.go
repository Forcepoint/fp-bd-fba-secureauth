package connectors

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"secureauth/internal/structs"
	"strings"
)

func SendToFBA(event structs.Event, conf *viper.Viper) {

	// Convert to Json
	jsonData, err := json.Marshal(event)
	if err != nil {
		logrus.Warn("An error occurred when converting the event data to json.")
		return
	}

	// Send events
	url := fmt.Sprintf("%s:%s/event", conf.GetString("fba_endpoint"), conf.GetString("fba_port"))
	success := Send("POST", url, jsonData, nil, conf)

	// Handle response
	if !success {
		logrus.Warn("The request to the FBA endpoint was not successful.")
		return
	}

	// Check if user in event
	numEntities := len(event.Entities)
	if !(event.Entities[numEntities - 1].Role == "User") {
		logrus.Info("No user present in event to be monitored.")
		return
	}

	// Set user to monitored
	user := event.Entities[numEntities - 1].Entities[0]
	success = monitorUser(user, conf)
	if !success {
		logrus.Warn("User could not be set to monitored.")
		return
	}

}

func monitorUser(user string, conf *viper.Viper) bool {

	// Check if monitored
	monitored := checkMonitored(user, conf)
	if monitored {
		return true
	}
	return setMonitored(user, conf)

}

func checkMonitored(user string, conf *viper.Viper) bool {

	// Build request to endpoint
	url := fmt.Sprintf("%s:%s/v1/entity/list/monitored", conf.GetString("rose_endpoint"), conf.GetString("rose_port"))
	respStruct := new(structs.MonitoredEntities)
	success := Send("GET", url, nil, respStruct, conf)

	// Handle response
	if !success {
		logrus.Warn("Unable to retrieve data from users endpoint.")
		return false
	}

	// Check if user is in response
	return checkInEntitiesArray(user, *respStruct)

}

func setMonitored(user string, conf *viper.Viper) bool {

	// Exit user does not exist
	if user == "" {
		return false
	}

	// Generate data and convert to json
	data := structs.MonitoredEntity{
		Name:  "Monitored Entity",
		Value: "TRUE",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		logrus.Warn("An error occurred when converting the event data to json.")
		return false
	}

	// Build endpoint for retrieving monitored users
	url := fmt.Sprintf("%s:%s/reference/actor/%s/attribute/boolean", conf.GetString("mds_endpoint"), conf.GetString("mds_port"), strings.ToLower(user))
	success := Send("POST", url, jsonData, nil, conf)

	return success

}

func checkInEntitiesArray(user string, array structs.MonitoredEntities) bool {

	for _, entity := range array.Entities {
		if strings.ToLower(entity.ActorID) == strings.ToLower(user) {
			return true
		}
	}
	return false

}