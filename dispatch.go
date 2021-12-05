package nod

func dispatch(msgType MessageType, payload interface{}, topic string) {
	for out, hnd := range handlers {
		if disabledOutputs[out] {
			continue
		}
		hnd.Handle(msgType, payload, topic)
	}
}
