package logs

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func InitLogs(conf *viper.Viper) {

	// Setup logger.
	file, err := os.OpenFile("events.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logrus.SetReportCaller(true)
	logrus.SetOutput(file)
	logrus.SetFormatter(&nested.Formatter{})

	// Set log level
	logrus.SetLevel(getLevel(conf.GetString("log_level")))

}

func getLevel(level string) logrus.Level {
	switch level {
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}