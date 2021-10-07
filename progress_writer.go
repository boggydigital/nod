package nod

import "io"

type TotalProgressWriter interface {
	io.Writer
	Progress(uint64)
	Total(uint64)
}

type totalProgress struct {
	topics  []string
	total   uint64
	current uint64
}

func (tp *totalProgress) Write(bytes []byte) (int, error) {
	tp.Progress(tp.current + uint64(len(bytes)))
	return len(bytes), nil
}

func (tp *totalProgress) Progress(current uint64) {
	tp.current = current
	dispatch(MsgProgress, tp.current, tp.topics...)
}

func (tp *totalProgress) Total(total uint64) {
	tp.total = total
	dispatch(MsgTotal, tp.total, tp.topics...)
}

func TotalProgress(topics ...string) *totalProgress {
	return &totalProgress{
		topics: topics,
	}
}
