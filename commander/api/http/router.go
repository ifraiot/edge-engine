package http

import (
	"edgeengine/commander/handlers"

	"github.com/gin-gonic/gin"
)

type endpoint struct {
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	serviceHandler := handlers.NewServiceHandler("../../build-in-app.json")

	serviceEndpoint := NewServiceEndpoint(serviceHandler)

	serviceGroup := router.Group("available-applications")
	serviceGroup.GET("", serviceEndpoint.AvailableApplication)

	return router
}
