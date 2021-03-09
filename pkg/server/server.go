package server

import (
	"gin-api-template/pkg/config"
	"gin-api-template/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

type server struct {
	engine *gin.Engine
}

// Init initializes web server
func Init() {
	gin.SetMode(config.Config().GetString("gin.mode"))

	srv := newServer()
	srv.run()
}

// newServer generates web server
func newServer() server {
	srv := server{}

	srv.engine = gin.Default()

	// CORS setup
	corsConfig := cors.DefaultConfig()
	setCorsOrigins(&corsConfig, util.RetrieveEnv("CORS_ORIGINS", true))
	srv.engine.Use(cors.New(corsConfig))

	srv.setRoutes()

	return srv
}

// run starts web server
func (srv *server) run() {
	err := srv.engine.Run(":8080")
	if err != nil {
		log.Fatalf("Starting server was failed\n%v", err)
	}
}

// setCorsOrigins sets CORS origin name from comma separated string
func setCorsOrigins(config *cors.Config, originsString string) {
	origins := strings.Split(originsString, ",")

	if len(origins) == 1 && len(origins[0]) == 0 {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = origins
	}
}
