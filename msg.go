package utr

func msg(msgType MessageType, payload interface{}, topics ...string) {
	for _, lst := range listeners[msgType] {
		lst.Listen(msgType, payload, topics...)
	}
}

func Begin(topics ...string) {
	msg(MsgBegin, nil, topics...)
}

func End(topics ...string) {
	msg(MsgEnd, nil, topics...)
}

func Success(success bool, topics ...string) {
	msg(MsgSuccess, success, topics...)
}

func Error(err error, topics ...string) {
	msg(MsgError, err, topics...)
}

func Summary(sum map[string][]string, topics ...string) {
	msg(MsgSummary, sum, topics...)
}

func Total(total uint64, topics ...string) {
	msg(MsgTotal, total, topics...)
}

func Progress(upd uint64, topics ...string) {
	msg(MsgProgress, upd, topics...)
}

func Log(log string, topics ...string) {
	msg(MsgLog, log, topics...)
}

func Debug(dbg string, topics ...string) {
	msg(MsgDebug, dbg, topics...)
}
