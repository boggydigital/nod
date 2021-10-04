package utter

import "fmt"

func EnableStdOut() {
	HandleFunc(&stdOutPresenter{}, StdOutTypes()...)
}

type stdOutPresenter struct {
}

func (sop *stdOutPresenter) Listen(msgType MessageType, payload interface{}, topics ...string) {
	switch msgType {
	case Begin:
		fmt.Println("begin", topics)
	case End:
		fmt.Println("end", topics)
	case Summary:
		fmt.Println("SUMMARY:")
		fmt.Println(payload)
	case Success:
		fmt.Println("SUCCESS:", payload)
	}
}
