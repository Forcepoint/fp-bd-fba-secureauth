package logs

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"regexp"
	"secureauth/internal/structs"
	"strings"
)

func HandleLogs(conf *viper.Viper, channel chan structs.Event) {

	// Open the log file.
	confDir := conf.GetString("realm_dir")
	configs := GetRealms(confDir)
	for _, realm := range configs {
		go ReadLog(realm, channel)
	}

}

func GetRealms(dir string) []*viper.Viper {

	// Get all realms files
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		logrus.Fatal("Unable to retrieve log file configurations.")
	}

	// Load in realms from config file.
	realms := make([]*viper.Viper, 0)
	for _, file := range files {
		fileDetails := strings.Split(file.Name(), ".")
		config := viper.New()
		config.SetConfigName(fileDetails[0])
		config.AddConfigPath(dir)
		err = config.ReadInConfig()
		if err != nil {
			logrus.Fatal(err)
		}
		realms = append(realms, config)
	}

	return realms

}

func ReadLog(realm *viper.Viper, channel chan structs.Event) {

	// Get values from realm
	time := realm.GetString("latest_time")
	logfile := realm.GetString("file")
	realmName := realm.GetString("realm")
	application := realm.GetString("application")

	// Open and tail log file
	t, err := tail.TailFile(logfile, tail.Config{Follow: true, Poll: true})
	if err != nil {
		return
	}

	// Read through the log file
	for line := range t.Lines {
		t := line.Text

		re := regexp.MustCompile(`(\w*)="(\W*|\w*|\S*)"`)
		values := re.FindAllStringSubmatch(t, -1)
		if values == nil {
			continue
		}

		if structs.EventTypes[values[2][2]] != "" && values[3][2] > time {
			logVal := structs.CreateEvent(values, application, realmName)
			time = logVal.Timestamp
			realm.Set("latest_time", time)
			realm.WriteConfig()
			logrus.Info(fmt.Sprintf("Event type '%s' logged.", logVal.Subject))
			channel <- logVal
		}

	}

}