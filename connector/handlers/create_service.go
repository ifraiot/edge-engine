package handlers

import (
	"context"
	"edgeengine/connector"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	docker "github.com/docker/docker/client"
	"github.com/ifrasoft/logger"
)

var _ connector.Handler = (*CreateServiceHandler)(nil)

type CreateServiceParams struct {
	Image string `json:"image"`
	Name  string `json:"name"`
	From  string `json:"from"`
	Envs  []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"envs"`
}

type CreateServiceHandler struct {
	cli *docker.Client
	lg  logger.Logger
}

func NewCreateServiceHandler(cli *docker.Client, lg logger.Logger) *CreateServiceHandler {
	return &CreateServiceHandler{cli: cli, lg: lg}
}

func (h *CreateServiceHandler) CommandName() string {

	return "create_service"
}

func (h *CreateServiceHandler) Handle(command connector.Command) ([]byte, error) {

	params := CreateServiceParams{}
	err := command.Bind(&params)
	if err != nil {
		return []byte(""), err
	}

	ctx := context.Background()
	reader, err := h.cli.ImagePull(ctx, params.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := h.cli.ContainerCreate(ctx, &container.Config{

		Image: params.Image,
		// Cmd:   []string{"echo", "hello world"},

	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := h.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := h.cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := h.cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

	return []byte(""), nil
}
