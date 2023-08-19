package http

import (
	"edgeengine/commander/handlers"

	"github.com/gin-gonic/gin"
)

type endpoint struct {
}

func NewRouter(uiEnabled bool) *gin.Engine {
	router := gin.Default()

	serviceHandler := handlers.NewServiceHandler("../../build-in-app.json")

	serviceEndpoint := NewServiceEndpoint(serviceHandler)

	serviceGroup := router.Group("/api")
	serviceGroup.GET("/available-applications", serviceEndpoint.AvailableApplication)

	if uiEnabled {
		router.Static("/", "../../ui/dist")
	}

	return router
}
