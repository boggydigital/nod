package nod

type MessageType int

const (
	//MsgStart signals start of an activity. No payload
	MsgStart MessageType = iota
	//MsgEnd signals completion of an activity. No payload
	MsgEnd
	//MsgSuccess passes binary result of an activity. Payload type: bool
	MsgSuccess
	//MsgError provides not-fatal error that happened during activity execution. Payload: error
	MsgError
	//MsgSummary provides map of categorized results. Payload: map[string][]string
	MsgSummary
	//MsgTotal sets expected total value of an activity progress. Payload: uint64
	MsgTotal
	//MsgProgress updates current value of an activity progress. Payload: uint64
	MsgProgress
	//MsgLog sends a value useful for logging. Payload: string
	MsgLog
	//MsgDebug sends a value useful for debugging. Payload: string
	MsgDebug
)

func StdOutTypes() []MessageType {
	return []MessageType{
		MsgStart,
		MsgEnd,
		MsgSuccess,
		MsgError,
		MsgSummary,
		MsgTotal,
		MsgProgress,
	}
}

func LogTypes() []MessageType {
	return append(StdOutTypes(), MsgLog)
}

func DebugTypes() []MessageType {
	return append(LogTypes(), MsgDebug)
}
