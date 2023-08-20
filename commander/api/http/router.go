package http

import (
	"edgeengine/commander/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type endpoint struct {
}

func NewRouter(uiEnabled bool) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	serviceHandler := handlers.NewServiceHandler("../../build-in-app.json")

	serviceEndpoint := NewServiceEndpoint(serviceHandler)

	serviceGroup := router.Group("/api")
	serviceGroup.GET("/available-applications", serviceEndpoint.AvailableApplication)

	if uiEnabled {
		// router.Static("/", "../../ui/dist")
	}

	return router
}
