package handlers

import "edgeengine/connector"

var _ connector.Handler = (*CreateServiceHandler)(nil)

type ReadServiceStatusHandler struct{}

func NewReadServiceStatusHandler() *CreateServiceHandler {
	return &CreateServiceHandler{}
}

func (h *ReadServiceStatusHandler) CommandName() string {
	return "inspect_service"
}

func (h *ReadServiceStatusHandler) Handle(command connector.Command) error {
	return nil
}
