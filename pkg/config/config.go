package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var conf *viper.Viper

// Init initializes config
func Init(fileName string) {
	log.Info("Start to load config")

	conf = viper.New()

	conf.SetConfigFile("YAML")
	conf.SetConfigName(fileName)

	conf.AddConfigPath("configs/environments/")
	conf.AddConfigPath("/run/secrets/")

	err := conf.ReadInConfig()
	if err != nil {
		log.Fatalf("Loading config was failed\n%v", err)
	}

	log.Info("Loading config has completed")
}

// Config returns config
func Config() *viper.Viper {
	return conf
}
