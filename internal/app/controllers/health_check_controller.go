package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheckController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
