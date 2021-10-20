package nod

import (
	"log"
	"strings"
)

func EnableStdOutRepeater() {
	sor := &stdOutRepeater{}
	HandleFunc(sor, DebugTypes()...)
}

type stdOutRepeater struct {
}

func (sor *stdOutRepeater) Handle(msgType MessageType, payload interface{}, topic string) {
	log.Println(strings.ToUpper(msgType.String()), topic, payload)
}
