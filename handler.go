package nod

type Handler interface {
	Handle(msgType MessageType, payload interface{}, topic string)
}

var handlers = make(map[MessageType][]Handler)

func HandleFunc(handler Handler, msgTypes ...MessageType) {
	for _, msgType := range msgTypes {
		handlers[msgType] = append(handlers[msgType], handler)
	}
}
