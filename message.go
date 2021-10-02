package utter

type MessageType int

const (
	//Begin signals start of an activity. No payload
	Begin MessageType = iota
	//End signals completion of an activity. No payload
	End
	//Success passes binary result of an activity. Payload type: bool
	Success
	//Error provides not-fatal error that happened during activity execution. Payload: error
	Error
	//Summary provides map of categorized results. Payload: map[string][]string
	Summary
	//Total sets expected total value of an activity progress. Payload: uint64
	Total
	//Progress updates current value of an activity progress. Payload: uint64
	Progress
	//Log sends a value useful for logging. Payload: string
	Log
	//Debug sends a value useful for debugging. Payload: string
	Debug
)

func StdOut() []MessageType {
	return []MessageType{
		Begin,
		End,
		Success,
		Error,
		Summary,
		Total,
		Progress,
	}
}

func StdErr() []MessageType {
	return append(StdOut(), Log)
}

func StdDbg() []MessageType {
	return append(StdErr(), Debug)
}
