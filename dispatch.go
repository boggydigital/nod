package nod

func dispatch(msgType MessageType, payload interface{}, topic string) {
	for _, h := range handlers[msgType] {
		h.Handle(msgType, payload, topic)
	}
}
