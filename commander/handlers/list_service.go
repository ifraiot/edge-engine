package handlers

import (
	"context"
	"edgeengine/commander"
	"encoding/json"

	"github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"
	"github.com/ifrasoft/logger"
)

var _ commander.Handler = (*ListServiceHandler)(nil)

type ListServiceParams struct {
	Image string `json:"image"`
	Name  string `json:"name"`
	From  string `json:"from"`
	Envs  []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"envs"`
}

type ListServiceHandler struct {
	cli *docker.Client
	lg  logger.Logger
}

func NewListServiceHandler(cli *docker.Client, lg logger.Logger) *ListServiceHandler {
	return &ListServiceHandler{cli: cli, lg: lg}
}

func (h *ListServiceHandler) CommandName() string {
	return "list_service"
}

func (h *ListServiceHandler) Handle(command commander.Command) ([]byte, error) {

	params := CreateServiceParams{}
	err := command.Bind(&params)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	containers, err := h.cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(containers)
	if err != nil {
		return nil, err
	}

	return data, err
}
