package utter

func Msg(msgType MessageType, payload interface{}, topics ...string) {
	for _, lst := range listeners[msgType] {
		lst.Listen(msgType, payload, topics...)
	}
}
