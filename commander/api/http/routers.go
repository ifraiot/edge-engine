package http

import (
	"edgeengine/commander"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceEndpoint struct {
	commanderService commander.CommanderService
}

func NewServiceEndpoint(commanderService commander.CommanderService) *ServiceEndpoint {
	return &ServiceEndpoint{
		commanderService: commanderService,
	}
}

type InstallApplicationRequest struct {
	AppId   string             `json:"app_id"`
	Configs []commander.Config `json:"configs"`
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

	err := h.commanderService.UninstallApplication(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

}

func (h *ServiceEndpoint) InstalledApplication(c *gin.Context) {

	installedApplications, err := h.commanderService.InstalledApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	availableApplications, err := h.commanderService.AvailableApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	var availableApplicationsMaps = make(map[string]commander.AvailableApplication, len(availableApplications))
	for _, application := range availableApplications {
		availableApplicationsMaps[application.ID] = application
	}

	for i, installedApplication := range installedApplications {

		status, _ := h.commanderService.ApplicationStatus(installedApplication.ID)

		installedApplications[i].Application = availableApplicationsMaps[installedApplication.AppID]
		installedApplications[i].Status = status
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

	availableApplications, err := h.commanderService.AvailableApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	var availableApplication commander.AvailableApplication
	for _, app := range availableApplications {
		if app.ID == installApplicationRequest.AppId {
			availableApplication = app
		}
	}

	err = h.commanderService.InstallApplication(availableApplication, installApplicationRequest.Configs)
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
	availableApplications, err := h.commanderService.AvailableApplications()
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
