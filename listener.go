package utr

type Listener interface {
	Listen(msgType MessageType, payload interface{}, topics ...string)
}

var listeners = make(map[MessageType][]Listener)

func HandleFunc(listener Listener, msgTypes ...MessageType) {
	for _, msgType := range msgTypes {
		listeners[msgType] = append(listeners[msgType], listener)
	}
}
