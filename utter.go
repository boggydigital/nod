package utter

func Utter(msgType MessageType, payload interface{}, path ...string) {
	for _, lst := range listeners[msgType] {
		go lst.Listen(msgType, payload, path...)
	}
}
