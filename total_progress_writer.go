package nod

import (
	"io"
)

type TotalProgressWriter interface {
	Total(uint64)
	Current(uint64)
	Progress(uint64)
	TotalInt(int)
	CurrentInt(int)
	ProgressInt(int)
	Increment()
	io.Writer
	ActLogCloser
}

type totalProgress struct {
	activity
	current uint64
	total   uint64
}

func (tp *totalProgress) Write(bytes []byte) (int, error) {
	tp.Progress(uint64(len(bytes)))
	return len(bytes), nil
}

func (tp *totalProgress) Total(total uint64) {
	tp.total = total
	dispatch(MsgTotal, tp.total, tp.topic)
}

func (tp *totalProgress) TotalInt(total int) {
	tp.Total(uint64(total))
}

func (tp *totalProgress) Current(current uint64) {
	tp.current = current
	dispatch(MsgCurrent, tp.current, tp.topic)
}

func (tp *totalProgress) CurrentInt(current int) {
	tp.Current(uint64(current))
}

func (tp *totalProgress) Progress(value uint64) {
	tp.current += value
	dispatch(MsgCurrent, tp.current, tp.topic)
}

func (tp *totalProgress) ProgressInt(value int) {
	tp.Progress(uint64(value))
}

func (tp *totalProgress) Increment() {
	tp.Progress(1)
}

func NewProgress(format string, d ...interface{}) TotalProgressWriter {
	return &totalProgress{
		activity: *Begin(format, d...),
	}
}
