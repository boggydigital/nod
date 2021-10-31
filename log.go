package nod

import "fmt"

func Log(format string, d ...interface{}) {
	dispatch(MsgLog, nil, fmt.Sprintf(format, d...))
}
