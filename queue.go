package nod

import "fmt"

type message struct {
	msgType MessageType
	payload interface{}
	topic   string
}

type Queue interface {
	Flush()
	Discard()
}

type queue struct {
	activity
	messages []message
}

func (q *queue) Discard() {
	q.messages = nil
	q.active = false
}

func (q *queue) Flush() {
	for _, m := range q.messages {
		dispatch(m.msgType, m.payload, m.topic)
	}
	q.Discard()
}

func QueueBegin(format string, d ...interface{}) *queue {
	topic := fmt.Sprintf(format, d...)
	messages := make([]message, 1)
	messages[0] = message{
		msgType: MsgBegin,
		topic:   topic,
	}
	return &queue{
		activity: activity{
			topic:  topic,
			active: true,
		},
		messages: messages,
	}
}

func (q *queue) End() {
	if q.active {
		q.messages = append(q.messages, message{
			msgType: MsgEnd,
			payload: nil,
			topic:   q.topic,
		})
	}
	q.active = false
}

func (q *queue) EndWithResult(format string, d ...interface{}) {
	if q.active {
		q.messages = append(q.messages, message{
			msgType: MsgResult,
			payload: fmt.Sprintf(format, d...),
			topic:   q.topic,
		})
		q.End()
	}
}

func (q *queue) Error(err error) {
	if q.active {
		q.messages = append(q.messages, message{
			msgType: MsgError,
			payload: err,
			topic:   q.topic,
		})
	}
}

func (q *queue) EndWithError(err error) error {
	if q.active {
		q.Error(err)
		q.End()
		return err
	}
	return nil
}

func (q *queue) EndWithSummary(summary headingSections) {
	if q.active {
		q.messages = append(q.messages, message{
			msgType: MsgSummary,
			payload: summary,
			topic:   q.topic,
		})
	}
}

func (q *queue) Log(format string, d ...interface{}) {
	if q.active {
		q.messages = append(q.messages, message{
			msgType: MsgLog,
			payload: fmt.Sprintf(format, d...),
			topic:   q.topic,
		})
	}
}
