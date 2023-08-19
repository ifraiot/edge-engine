package handlers

import (
	"context"
	"edgeengine/connector"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	docker "github.com/docker/docker/client"
	"github.com/ifrasoft/logger"
)

var _ connector.Handler = (*TerminateServiceStatusHandler)(nil)

type TerminateServiceParams struct {
	NameOrID string `json:"name"`
}

type TerminateServiceStatusHandler struct {
	cli *docker.Client
	lg  logger.Logger
}

func NewTerminateServiceStatusHandler(cli *docker.Client, lg logger.Logger) *TerminateServiceStatusHandler {
	return &TerminateServiceStatusHandler{cli: cli, lg: lg}
}

func (h *TerminateServiceStatusHandler) CommandName() string {
	return "terminate_service"
}

func (h *TerminateServiceStatusHandler) Handle(command connector.Command) ([]byte, error) {

	params := TerminateServiceParams{}
	err := command.Bind(&params)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	// Stop the container
	if err := h.cli.ContainerStop(ctx, params.NameOrID, container.StopOptions{}); err != nil {
		fmt.Printf("Failed to stop container: %s\n", err)
		return nil, err
	}

	fmt.Printf("Container %s stopped\n", params.NameOrID)

	// Remove the container
	if err := h.cli.ContainerRemove(ctx, params.NameOrID, types.ContainerRemoveOptions{}); err != nil {
		fmt.Printf("Failed to remove container: %s\n", err)
		return nil, err
	}
	fmt.Printf("Container %s removed\n", params.NameOrID)

	return nil, nil
}
