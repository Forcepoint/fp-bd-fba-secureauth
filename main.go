package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"secureauth/internal/config"
	"secureauth/internal/connectors"
	"secureauth/internal/logs"
	"secureauth/internal/realms"
	"secureauth/internal/structs"
)

var conf *viper.Viper

func init() {

	// Initialise config
	conf = config.InitConfig()

	// Initialise logging
	logs.InitLogs(conf)

	// Retrieve realms
	realms.GetRealms(conf)

}

func main() {

	// Retrieve events
	eventsChannel := make(chan structs.Event)
	go logs.HandleLogs(conf, eventsChannel)

	// Handle events
	for {
		event := <- eventsChannel
		logrus.Info(fmt.Sprintf("New Event Received: %s", event.SourceEventID))

		go connectors.SendToFBA(event, conf)
	}
}