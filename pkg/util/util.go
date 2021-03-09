package util

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func RetrieveEnv(name string, allowEmpty bool) (env string) {
	env = os.Getenv(name)
	if env == "" && !allowEmpty {
		log.WithFields(log.Fields{
			"env_name": name,
		}).Fatal("Environment is not set")
	}

	return
}
