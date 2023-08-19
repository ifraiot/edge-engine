package http

import (
	"edgeengine/commander/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceEndpoint struct {
	serviceHandler handlers.ServiceHandler
}

func NewServiceEndpoint(serviceHandler handlers.ServiceHandler) *ServiceEndpoint {
	return &ServiceEndpoint{
		serviceHandler: serviceHandler,
	}
}

func (h *ServiceEndpoint) AvailableApplication(c *gin.Context) {
	availableApplications, err := h.serviceHandler.AvailableApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   availableApplications,
	})
}
