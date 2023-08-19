package handlers

import (
	"edgeengine/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ServiceHandler interface {
	AvailableApplications() (availableApplications []service.AvailableApplication, err error)
}

type serviceHandler struct {
	appPath string
}

type Configuration struct {
	Connectors  []service.AvailableApplication `json:"connectors"`
	Analyzers   []service.AvailableApplication `json:"analyzers"`
	Integrators []service.AvailableApplication `json:"integrators"`
}

func NewServiceHandler(appPath string) ServiceHandler {
	return &serviceHandler{appPath: appPath}
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

	availableApplications = append(availableApplications, config.Connectors...)
	availableApplications = append(availableApplications, config.Analyzers...)
	availableApplications = append(availableApplications, config.Integrators...)

	return

}
