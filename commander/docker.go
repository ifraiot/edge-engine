package commander

type CreateContainerParams struct {
	Image string `json:"image"`
	Name  string `json:"name"`
	From  string `json:"from"`
	Envs  []Env  `json:"envs"`
	Ports []Port `json:"ports"`
}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Port struct {
	Public    string `json:"public"`
	Container string `json:"container"`
}

type DockerAPI interface {
	CreateContainer(param CreateContainerParams) (string, error)
	DeleteContainer(containerId string) error
}
