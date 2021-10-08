package nod

func dispatch(msgType MessageType, payload interface{}, topics ...string) {
	for _, h := range handlers[msgType] {
		h.Handle(msgType, payload, topics...)
	}
}

func Start(topics ...string) {
	dispatch(MsgStart, nil, topics...)
}

func End(topics ...string) {
	dispatch(MsgEnd, nil, topics...)
}

func Result(res string, topics ...string) {
	dispatch(MsgResult, res, topics...)
}

func Error(err error, topics ...string) {
	dispatch(MsgError, err, topics...)
}

func Fatal(err error, topics ...string) error {
	Error(err, topics...)
	End(topics...)
	return err
}

func Summary(sum map[string][]string, topics ...string) {
	dispatch(MsgSummary, sum, topics...)
}

func Total(total uint64, topics ...string) {
	dispatch(MsgTotal, total, topics...)
}

func Progress(upd uint64, topics ...string) {
	dispatch(MsgProgress, upd, topics...)
}

func Log(log string, topics ...string) {
	dispatch(MsgLog, log, topics...)
}

func Debug(dbg string, topics ...string) {
	dispatch(MsgDebug, dbg, topics...)
}
