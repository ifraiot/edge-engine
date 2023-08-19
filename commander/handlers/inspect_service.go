package handlers

import "edgeengine/commander"

var _ commander.Handler = (*CreateServiceHandler)(nil)

type ReadServiceStatusHandler struct{}

func NewReadServiceStatusHandler() *CreateServiceHandler {
	return &CreateServiceHandler{}
}

func (h *ReadServiceStatusHandler) CommandName() string {
	return "inspect_service"
}

func (h *ReadServiceStatusHandler) Handle(command commander.Command) error {
	return nil
}
