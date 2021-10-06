package nod

type writer struct {
	topics  []string
	current uint64
}

func (w *writer) Write(bytes []byte) (int, error) {
	w.current = w.current + uint64(len(bytes))
	dispatch(MsgProgress, w.current, w.topics...)
	return len(bytes), nil
}

func ProgressWriter(topics ...string) *writer {
	return &writer{
		topics: topics,
	}
}
