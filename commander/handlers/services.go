package handlers

import (
	"edgeengine/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gorm.io/gorm"
)

type ServiceHandler interface {
	AvailableApplications() (availableApplications []service.AvailableApplication, err error)
	InstallApplication(appId string, configs []service.Config) error
}

type serviceHandler struct {
	db      *gorm.DB
	appPath string
}

type Configuration struct {
	Connectors  []service.AvailableApplication `json:"connectors"`
	Analyzers   []service.AvailableApplication `json:"analyzers"`
	Integrators []service.AvailableApplication `json:"integrators"`
}

func NewServiceHandler(db *gorm.DB, appPath string) ServiceHandler {
	return &serviceHandler{
		db:      db,
		appPath: appPath,
	}
}

func (h *serviceHandler) InstallApplication(appId string, configs []service.Config) error {

	jsonData, err := json.Marshal(configs)
	if err != nil {
		return err
	}

	err = h.db.Create(&service.InstalledApplication{
		AppID:  appId,
		Config: string(jsonData),
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (h *serviceHandler) InstalledApplications() error {
	return nil
}

func (h *serviceHandler) AvailableApplications() (availableApplications []service.AvailableApplication, err error) {

	file, err := os.Open(h.appPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Read the content of the file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Create a Configuration struct to hold the JSON data
	var config Configuration

	// Unmarshal JSON data into the struct
	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	availableApplications = append(availableApplications, func(availableApplications []service.AvailableApplication) []service.AvailableApplication {
		for i, _ := range availableApplications {
			availableApplications[i].Type = "connector"
		}
		return availableApplications
	}(config.Connectors)...)

	availableApplications = append(availableApplications, func(availableApplications []service.AvailableApplication) []service.AvailableApplication {
		for i, _ := range availableApplications {
			availableApplications[i].Type = "analyzer"
		}
		return availableApplications
	}(config.Analyzers)...)

	availableApplications = append(availableApplications, func(availableApplications []service.AvailableApplication) []service.AvailableApplication {
		for i, _ := range availableApplications {
			availableApplications[i].Type = "integrator"
		}
		return availableApplications
	}(config.Integrators)...)

	return

}
