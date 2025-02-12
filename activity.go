package nod

import "fmt"

type ActCloser interface {
	Log(string, ...interface{})
	Error(error)
	EndWithResult(string, ...interface{})
	EndWithSummary(string, map[string][]string)
}
type activity struct {
	topic  string
	active bool
}

func Begin(format string, d ...interface{}) *activity {
	topic := fmt.Sprintf(format, d...)
	dispatch(MsgBegin, nil, topic)
	return &activity{
		topic:  topic,
		active: true,
	}
}

func (a *activity) end() {
	if a.active {
		dispatch(MsgEnd, nil, a.topic)
		a.active = false
	}
}

func (a *activity) EndWithResult(format string, d ...interface{}) {
	if a.active {
		result := fmt.Sprintf(format, d...)
		dispatch(MsgResult, result, a.topic)
		a.end()
	}
}

func (a *activity) Error(err error) {
	if a.active {
		dispatch(MsgError, err, a.topic)
	}
}

func (a *activity) EndWithSummary(heading string, sections map[string][]string) {
	if a.active {
		dispatch(MsgSummary, headingSections{heading: heading, sections: sections}, a.topic)
		a.end()
	}
}

func (a *activity) Log(format string, d ...interface{}) {
	msg := fmt.Sprintf(format, d...)
	dispatch(MsgLog, msg, a.topic)
}
