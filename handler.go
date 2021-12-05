package nod

type Handler interface {
	Handle(msgType MessageType, payload interface{}, topic string)
	Close() error
}

var handlers = make(map[string]Handler)

func HandleFunc(handler Handler, output string) {
	handlers[output] = handler
}
