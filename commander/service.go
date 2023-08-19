package commander

type Handler interface {
	CommandName() string
	Handle(comMsg Command) ([]byte, error)
}
