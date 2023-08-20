package service

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
}

type InstalledApplication struct {
	ID          string               `json:"id"`
	AppID       string               `json:"app_id"`
	Config      string               `json:"config"`
	Application AvailableApplication `json:"application" gorm:"-"`
}
