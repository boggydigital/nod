package utr

import "fmt"

func EnableStdOut() {
	HandleFunc(&stdOutPresenter{}, StdOutTypes()...)
}

type stdOutPresenter struct {
}

func (sop *stdOutPresenter) Listen(msgType MessageType, payload interface{}, topics ...string) {
	switch msgType {
	case MsgBegin:
		fmt.Println("begin", topics)
	case MsgEnd:
		fmt.Println("end", topics)
	case MsgSummary:
		fmt.Println("SUMMARY:")
		fmt.Println(payload)
	case MsgSuccess:
		fmt.Println("SUCCESS:", payload)
	}
}
