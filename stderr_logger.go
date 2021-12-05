package nod

import (
	"log"
	"strings"
)

func EnableStdErrLogger() {
	sol := &stdErrLogger{}
	HandleFunc(sol, StdErr)
}

type stdErrLogger struct {
}

func (sel *stdErrLogger) Close() error {
	return nil
}

func (sel *stdErrLogger) Handle(msgType MessageType, payload interface{}, topic string) {
	if payload != nil {
		log.Println(strings.ToUpper(msgType.String()), topic, payload)
	} else {
		log.Println(strings.ToUpper(msgType.String()), topic)
	}

}
