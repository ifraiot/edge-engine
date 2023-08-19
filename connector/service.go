package connector

type Handler interface {
	CommandName() string
	Handle(comMsg Command) ([]byte, error)
}
