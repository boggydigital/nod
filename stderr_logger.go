package nod

import (
	"log"
	"strings"
)

func EnableStdErrLogger() {
	sol := &stdErrLogger{}
	HandleFunc(sol, LogTypes()...)
}

type stdErrLogger struct {
}

func (sol *stdErrLogger) Handle(msgType MessageType, payload interface{}, topic string) {
	log.Println(strings.ToUpper(msgType.String()), topic, payload)
}
