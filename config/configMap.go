package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type GlobalMap struct {
	DatabaseUrl string
	ServerPort  int
}

var GlobalMapEntity GlobalMap

func init() {
	viper.SetConfigType("yaml")

	if os.Getenv("DEPLOY_MODE") == "dev" {
		viper.AddConfigPath("config/dev")
	} else {
		viper.AddConfigPath("config/prod")
	}

	viper.SetConfigName("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panicln(err.Error())
		panic(err)
	} else {
		GlobalMapEntity.DatabaseUrl = viper.GetString("datasource.url")
		GlobalMapEntity.ServerPort = viper.GetInt("server.port")
	}
}
