package nod

import "fmt"

func Error(err error) {
	dispatch(MsgError, nil, err.Error())
}

func ErrorStr(format string, d ...interface{}) string {
	msg := fmt.Sprintf(format, d)
	dispatch(MsgError, nil, msg)
	return msg
}
