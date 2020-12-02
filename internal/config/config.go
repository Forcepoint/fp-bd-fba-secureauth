package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {

	// Load in configuration file.
	conf := viper.New()
	conf.SetConfigName("config")
	conf.AddConfigPath("./config")
	err := conf.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal Error. Could not locate config file: %s \n", err))
	}

	return conf

}