package docker

import (
	"context"
	"edgeengine/commander"

	"encoding/json"

	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	docker "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"

	"github.com/ifrasoft/logger"
)

const networkName = "ifra-network"

type dockerAPI struct {
	cli *docker.Client
	lg  logger.Logger
}

func NewDockerAPI(cli *docker.Client, lg logger.Logger) commander.DockerAPI {
	return &dockerAPI{cli: cli, lg: lg}
}

func (h *dockerAPI) ContainerInfo(containerID string) (types.ContainerJSON, error) {
	containerInfo, err := h.cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return containerInfo, nil
}

func (h *dockerAPI) DeleteContainer(containerId string) error {
	err := h.cli.ContainerStop(context.Background(), containerId, container.StopOptions{})
	if err != nil {
		return err
	}

	err = h.cli.ContainerRemove(context.Background(), containerId, types.ContainerRemoveOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (h *dockerAPI) CreateContainer(param commander.CreateContainerParams) (string, error) {

	ctx := context.Background()
	reader, err := h.cli.ImagePull(ctx, param.Image, types.ImagePullOptions{})
	if err != nil {
		return "", err
	}

	io.Copy(os.Stdout, reader)

	exists, err := overlayNetworkExists(h.cli, networkName)
	if err != nil {
		return "", err
	}

	if exists {
		fmt.Println("Overlay network already exists.")
	} else {
		err := createOverlayNetwork(h.cli, networkName)
		if err != nil {
			panic(err)
		}
		fmt.Println("Overlay network created.")
	}

	endpointConfig := &network.EndpointSettings{
		NetworkID: networkName,
	}

	hostConfig := &container.HostConfig{
		// Use the HostConfig field to set --add-host
		RestartPolicy: container.RestartPolicy{Name: "always"},
		ExtraHosts: []string{
			"host.docker.internal:host-gateway",
		},
		// Binds: []string{
		// 	"/path/to/host/synthesis.yml:/path/in/container/synthesis.yml",
		// },
		// PortBindings: portBinding(portMapping),
	}

	networkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			networkName: endpointConfig,
		},
	}

	// portMapping := make(map[string]string, len(param.Ports))
	// for _, port := range param.Ports {
	// 	portMapping[port.Container+"/tcp"] = port.Public
	// }

	strEnvs := []string{}
	for _, env := range param.Envs {
		strEnvs = append(strEnvs, fmt.Sprintf("%s=%s", env.Name, env.Value))
	}

	resp, err := h.cli.ContainerCreate(ctx, &container.Config{

		Env:   strEnvs,
		Image: param.Image,
	}, hostConfig, networkingConfig, nil, param.Name)
	if err != nil {
		return "", err
	}

	if err := h.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	// statusCh, errCh := h.cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	// select {
	// case err := <-errCh:
	// 	if err != nil {
	// 		return "", err
	// 	}
	// case <-statusCh:
	// }

	// out, err := h.cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(out)

	return resp.ID, nil
}

func (h *dockerAPI) ContainerList() ([]byte, error) {

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

func portBinding(portMap map[string]string) nat.PortMap {
	portBinding := make(nat.PortMap)
	for containerPort, hostPort := range portMap {
		port, _ := nat.NewPort("tcp", containerPort)
		portBinding[port] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: hostPort,
			},
		}
	}
	return portBinding
}

func overlayNetworkExists(cli *client.Client, networkName string) (bool, error) {
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return false, err
	}

	for _, network := range networks {
		if network.Name == networkName {
			return true, nil
		}
	}

	return false, nil
}

func createOverlayNetwork(cli *client.Client, networkName string) error {
	_, err := cli.NetworkCreate(context.Background(), networkName, types.NetworkCreate{})
	return err
}
