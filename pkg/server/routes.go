package server

import "gin-api-template/internal/app/controllers"

// setRoutes sets routing
func (srv *server) setRoutes() {
	engine := srv.engine

	engine.GET("/ping", controllers.HealthCheckController)
}
