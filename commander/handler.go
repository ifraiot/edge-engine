package commander

import (
	"encoding/json"
	"fmt"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/ifrasoft/logger"
)

type handlerSvc struct {
	edgeId   string
	logger   logger.Logger
	handlers map[string]Handler
}

func New(edgeId string, handlers map[string]Handler, logger logger.Logger) *handlerSvc {
	return &handlerSvc{
		edgeId:   edgeId,
		logger:   logger,
		handlers: handlers,
	}
}

func (h *handlerSvc) Handler(client MQTT.Client, msg MQTT.Message) {

	comMsgs := []Command{}
	err := json.Unmarshal(msg.Payload(), &comMsgs)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	edgeId, err := extractEdgeIDUsingSplit(msg.Topic())
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	if edgeId != h.edgeId {
		h.logger.Info("Not for me")
		return
	}

	for _, comMsg := range comMsgs {

		if handler, ok := h.handlers[comMsg.Command]; ok {

			h.logger.Info(handler.CommandName())

			resp, err := handler.Handle(comMsg)
			if err != nil {
				h.logger.Error(err.Error())

				requestPayload := err.Error()
				if token := client.Publish("edge/logs", 0, false, requestPayload); token.Wait() && token.Error() != nil {
					fmt.Println(token.Error())
				}
			} else {

				if token := client.Publish("edge/logs", 0, false, resp); token.Wait() && token.Error() != nil {
					fmt.Println(token.Error())
				}
			}

		} else {
			h.logger.Error("Command Not Support")
			return
		}

	}
}

func extractEdgeIDUsingSplit(input string) (string, error) {
	parts := strings.Split(input, "/")
	if len(parts) >= 3 {
		return parts[1], nil
	}
	return "", fmt.Errorf("no edge ID found")
}
