package nod

import "fmt"

type ActCloser interface {
	End()
	EndWithResult(string, ...interface{})
	EndWithError(error) error
	EndWithSummary(map[string][]string)
}

type ActLogger interface {
	Error(error)
	Log(string, ...interface{})
	Debug(string, ...interface{})
}

type ActLogCloser interface {
	ActLogger
	ActCloser
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

func (a *activity) End() {
	if a.active {
		dispatch(MsgEnd, nil, a.topic)
		a.active = false
	}
}

func (a *activity) EndWithResult(format string, d ...interface{}) {
	if a.active {
		result := fmt.Sprintf(format, d...)
		dispatch(MsgResult, result, a.topic)
		a.End()
	}
}

func (a *activity) Error(err error) {
	if a.active {
		dispatch(MsgError, err, a.topic)
	}
}

func (a *activity) EndWithError(err error) error {
	if a.active {
		a.Error(err)
		a.End()
		return err
	}
	return nil
}

func (a *activity) EndWithSummary(summary map[string][]string) {
	if a.active {
		dispatch(MsgSummary, summary, a.topic)
		a.End()
	}
}

func (a *activity) Log(format string, d ...interface{}) {
	msg := fmt.Sprintf(format, d...)
	dispatch(MsgLog, msg, a.topic)
}

func (a *activity) Debug(format string, d ...interface{}) {
	dbg := fmt.Sprintf(format, d...)
	dispatch(MsgDebug, dbg, a.topic)
}
