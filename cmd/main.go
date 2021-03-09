package main

import (
	"gin-api-template/pkg/config"
	"gin-api-template/pkg/database"
	"gin-api-template/pkg/server"
	"gin-api-template/pkg/util"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	config.Init(util.RetrieveEnv("APP_ENV", false))

	database.Init(false)
	defer database.Close()

	server.Init()
}
