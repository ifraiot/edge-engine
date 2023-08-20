package http

import (
	"edgeengine/commander/handlers"
	"edgeengine/service"
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

type InstallApplicationRequest struct {
	AppId   string           `json:"app_id"`
	Configs []service.Config `json:"configs"`
}

type UninstallApplicationRequest struct {
	Id string `json:"id"`
}

func (h *ServiceEndpoint) UninstalledApplication(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "id is required",
		})
		return
	}

	err := h.serviceHandler.UninstallApplication(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

}

func (h *ServiceEndpoint) InstalledApplication(c *gin.Context) {

	installedApplications, err := h.serviceHandler.InstalledApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	availableApplications, err := h.serviceHandler.AvailableApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	var availableApplicationMaps = make(map[string]service.AvailableApplication, len(availableApplications))
	for _, availableApplication := range availableApplications {
		availableApplicationMaps[availableApplication.ID] = availableApplication
	}

	for i, _ := range installedApplications {
		installedApplications[i].Application = availableApplicationMaps[availableApplications[i].ID]
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   installedApplications,
	})
}

func (h *ServiceEndpoint) InstallApplication(c *gin.Context) {

	installApplicationRequest := InstallApplicationRequest{}
	err := c.ShouldBind(&installApplicationRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err = h.serviceHandler.InstallApplication(installApplicationRequest.AppId, installApplicationRequest.Configs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
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
