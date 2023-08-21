package commander

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type CommanderService interface {
	AvailableApplications() (availableApplications []AvailableApplication, err error)
	InstallApplication(availableApplication AvailableApplication, configs []Config) error
	InstalledApplications() ([]InstalledApplication, error)
	UninstallApplication(id string) error
	ApplicationStatus(id string) (string, error)
}

type commanderService struct {
	db        *gorm.DB
	appPath   string
	dockerAPI DockerAPI
}

type Configuration struct {
	Connectors  []AvailableApplication `json:"connectors"`
	Analyzers   []AvailableApplication `json:"analyzers"`
	Integrators []AvailableApplication `json:"integrators"`
}

var _ CommanderService = (*commanderService)(nil)

func NewCommandService(
	db *gorm.DB,
	appPath string,
	dockerAPI DockerAPI,
) CommanderService {
	return &commanderService{
		db:        db,
		appPath:   appPath,
		dockerAPI: dockerAPI,
	}
}

func (h *commanderService) ApplicationStatus(id string) (string, error) {

	var installedApplication InstalledApplication
	err := h.db.Where("id = ?", id).First(&installedApplication).Error
	if err != nil {
		return "", err
	}

	containerInfo, err := h.dockerAPI.ContainerInfo(installedApplication.ContainerId)
	if err != nil {
		return "", err
	}

	return containerInfo.State.Status, nil
}

func (h *commanderService) UninstallApplication(id string) error {

	tx := h.db.Begin()

	var installedApplication = InstalledApplication{ID: id}

	err := tx.First(&installedApplication).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where("id = ?", id).Delete(&InstalledApplication{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = h.dockerAPI.DeleteContainer(installedApplication.ContainerId)
	if err != nil {
		fmt.Println("Worning: ", err.Error())
	}

	tx.Commit()

	return nil
}

func (h *commanderService) InstallApplication(availableApplication AvailableApplication, configs []Config) error {

	tx := h.db.Begin()
	jsonData, err := json.Marshal(configs)
	if err != nil {
		tx.Rollback()
		return err
	}

	id := uuid.New().String()
	contrainerId, err := h.dockerAPI.CreateContainer(CreateContainerParams{
		Image: availableApplication.Registry + ":" + availableApplication.Tag,
		Name:  fmt.Sprintf("EdgeEngine-%ss-%s", availableApplication.ID, id),
		From:  "edgeengine",
		Envs: func(configs []Config) []Env {
			envs := []Env{}
			for _, config := range configs {
				envs = append(envs, Env{
					Name:  config.Name,
					Value: config.Value,
				})
			}
			return envs
		}(configs),
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Create(&InstalledApplication{
		ID:          id,
		AppID:       availableApplication.ID,
		ContainerId: contrainerId,
		Config:      string(jsonData),
	}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (h *commanderService) InstalledApplications() ([]InstalledApplication, error) {

	var installedApplications []InstalledApplication
	err := h.db.Find(&installedApplications).Error
	if err != nil {
		return nil, err
	}

	return installedApplications, nil
}

func (h *commanderService) AvailableApplications() (availableApplications []AvailableApplication, err error) {

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

	availableApplications = append(availableApplications, func(availableApplications []AvailableApplication) []AvailableApplication {
		for i, _ := range availableApplications {
			availableApplications[i].Type = "connector"
		}
		return availableApplications
	}(config.Connectors)...)

	availableApplications = append(availableApplications, func(availableApplications []AvailableApplication) []AvailableApplication {
		for i, _ := range availableApplications {
			availableApplications[i].Type = "analyzer"
		}
		return availableApplications
	}(config.Analyzers)...)

	availableApplications = append(availableApplications, func(availableApplications []AvailableApplication) []AvailableApplication {
		for i, _ := range availableApplications {
			availableApplications[i].Type = "integrator"
		}
		return availableApplications
	}(config.Integrators)...)

	return

}
