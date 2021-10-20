package nod

type MessageType int

const (
	//MsgNone is a message type initial value. No payload
	MsgNone MessageType = iota
	//MsgBegin signals start of an activity. No payload
	MsgBegin
	//MsgError provides not-fatal error that happened during activity execution. Payload: error
	MsgError
	//MsgTotal sets expected total value of an activity progress. Payload: uint64
	MsgTotal
	//MsgCurrent updates current value of an activity progress. Payload: uint64
	MsgCurrent
	//MsgLog sends a value useful for logging. Payload: string
	MsgLog
	//MsgDebug sends a value useful for debugging. Payload: string
	MsgDebug
	//MsgResult provides result of an activity. Payload type: string
	MsgResult
	//MsgSummary provides map of categorized results. Payload: map[string][]string
	MsgSummary
	//MsgEnd signals completion of an activity. No payload
	MsgEnd
)

var messageTypeStrings = map[MessageType]string{
	MsgNone:    "none",
	MsgBegin:   "begin",
	MsgError:   "error",
	MsgTotal:   "total",
	MsgCurrent: "current",
	MsgLog:     "log",
	MsgDebug:   "debug",
	MsgResult:  "result",
	MsgSummary: "summary",
	MsgEnd:     "end",
}

func StdOutTypes() []MessageType {
	return []MessageType{
		MsgBegin,
		MsgEnd,
		MsgError,
		MsgTotal,
		MsgCurrent,
		MsgResult,
		MsgSummary,
	}
}

func LogTypes() []MessageType {
	return append(StdOutTypes(), MsgLog)
}

func DebugTypes() []MessageType {
	return append(LogTypes(), MsgDebug)
}

func (mt MessageType) String() string {
	return messageTypeStrings[mt]
}
