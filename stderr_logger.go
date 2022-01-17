package nod

import (
	"fmt"
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

	if skipLogging(msgType) {
		return
	}

	logLine := fmt.Sprintf("%s %s", strings.ToUpper(msgType.String()), topic)

	if payload != nil {
		logLine = fmt.Sprintf("%s %v", logLine, payload)
	}

	log.Println(logLine)
}
