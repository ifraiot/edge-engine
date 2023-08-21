package commander

import (
	"encoding/json"
)

type AvailableApplication struct {
	ID       string   `json:"id"`
	Label    string   `json:"label"`
	Type     string   `json:"type"`
	Registry string   `json:"registry"`
	Tag      string   `json:"tag"`
	Config   []Config `json:"config"`
}

type Config struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Example    string `json:"example"`
	IsRequired bool   `json:"is_required"`
	Value      string `json:"value"`
	Default    string `json:"default"`
}

type InstalledApplication struct {
	ID          string               `json:"id" gorm:"primarykey"`
	AppID       string               `json:"app_id"`
	Config      string               `json:"config"`
	Application AvailableApplication `json:"application" gorm:"-"`
	ContainerId string               `json:"container_id"`
	Status      string               `json:"status" gorm:"-"`
}

type Command struct {
	Command    string      `json:"command"`
	Sender     string      `json:"sender"`
	Parameters interface{} `json:"parameters"`
}

func (c Command) Bind(params interface{}) error {

	data, err := json.Marshal(c.Parameters)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, params)
}
