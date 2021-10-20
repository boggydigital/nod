package nod

type ActCloser interface {
	End()
	EndWithResult(string)
	EndWithError(error) error
	EndWithSummary(map[string][]string)
}

type ActLogger interface {
	Error(error)
	Log(string)
	Debug(string)
}

type ActLogCloser interface {
	ActLogger
	ActCloser
}

type activity struct {
	topic  string
	active bool
}

func Begin(topic string) *activity {
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

func (a *activity) EndWithResult(result string) {
	if a.active {
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

func (a *activity) Log(msg string) {
	dispatch(MsgLog, msg, a.topic)
}

func (a *activity) Debug(dbg string) {
	dispatch(MsgDebug, dbg, a.topic)
}
