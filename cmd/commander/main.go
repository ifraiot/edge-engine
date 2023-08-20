package main

import (
	"edgeengine/commander"
	commandHTTP "edgeengine/commander/api/http"
	"edgeengine/commander/handlers"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	docker "github.com/docker/docker/client"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/ifrasoft/logger"
	// Sqlite driver based on CGO
)

const baseTopic = "edges/%s/commands"

func main() {

	lg, _ := logger.New(os.Stdout, "info")

	edgeId := "edge1"

	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://broker.hivemq.com:1883") // Replace with your broker's address
	opts.SetClientID("edge-agent")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		lg.Error("mqtt connection is failed")
		panic(token.Error())
	}

	cli, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		lg.Error("Failed to create docker client")
		panic(err)
	}

	r := commandHTTP.NewRouter(true)

	go http.ListenAndServe(":8000", r)

	defer cli.Close()

	handlers := map[string]commander.Handler{
		"create_service":    handlers.NewCreateServiceHandler(cli, lg),
		"inspect_service":   handlers.NewReadServiceStatusHandler(),
		"list_service":      handlers.NewListServiceHandler(cli, lg),
		"terminate_service": handlers.NewTerminateServiceStatusHandler(cli, lg),
	}

	handlerService := commander.New(edgeId, handlers, lg)

	if token := client.Subscribe(fmt.Sprintf(baseTopic, edgeId), 0, handlerService.Handler); token.Wait() && token.Error() != nil {
		lg.Error("Failed to subscribe:" + token.Error().Error())
	}

	lg.Info("edge engine is running")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel

	client.Disconnect(250)

	lg.Info("edge engine is terminated")
}
