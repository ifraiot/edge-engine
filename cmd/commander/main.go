package main

import (
	"edgeengine/commander"

	httpapi "edgeengine/commander/api/http"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	dockerAPI "edgeengine/commander/docker"

	docker "github.com/docker/docker/client"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ifrasoft/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const baseTopic = "edges/%s/commands"

func main() {

	lg, err := logger.New(os.Stdout, "info")
	if err != nil {
		panic(err)
	}

	// edgeId := "edge1"

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

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&commander.InstalledApplication{})

	dockerAPI := dockerAPI.NewDockerAPI(cli, lg)
	commanderService := commander.NewCommandService(db, "../../build-in-app.json", dockerAPI)
	serviceEndpoint := httpapi.NewServiceEndpoint(commanderService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//HTTP API
	serviceGroup := router.Group("/api")
	serviceGroup.GET("/available-applications", serviceEndpoint.AvailableApplication)
	appGroup := serviceGroup.Group("/applications")
	appGroup.POST("", serviceEndpoint.InstallApplication)
	appGroup.GET("", serviceEndpoint.InstalledApplication)
	appGroup.DELETE(":id", serviceEndpoint.UninstalledApplication)
	go http.ListenAndServe(":8000", router)
	defer cli.Close()

	// if token := client.Subscribe(fmt.Sprintf(baseTopic, edgeId), 0, handlerService.Handler); token.Wait() && token.Error() != nil {
	// 	lg.Error("Failed to subscribe:" + token.Error().Error())
	// }

	lg.Info("edge engine is running")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel

	client.Disconnect(250)

	lg.Info("edge engine is terminated")
}
