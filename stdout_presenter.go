package nod

import (
	"fmt"
	"path"
	"strings"
)

func EnableStdOut() {
	sop := &stdOutPresenter{
		topicTotals:   make(map[string]uint64),
		topicPercents: make(map[string]string),
	}
	HandleFunc(sop, StdOutTypes()...)
}

type stdOutPresenter struct {
	topicTotals   map[string]uint64
	topicPercents map[string]string
	startedTopics []string
}

func (sop *stdOutPresenter) Handle(msgType MessageType, payload interface{}, topics ...string) {
	switch msgType {
	case MsgStart:
		sop.handleBegin(topics...)
	case MsgEnd:
		sop.handleEnd(topics...)
	case MsgTotal:
		if total, ok := payload.(uint64); ok {
			sop.handleTotal(total, topics...)
		}
	case MsgProgress:
		if upd, ok := payload.(uint64); ok {
			sop.handleProgress(upd, topics...)
		}
	case MsgResult:
		if res, ok := payload.(string); ok {
			sop.handleResult(res, topics...)
		}
	case MsgSummary:
		if sum, ok := payload.(map[string][]string); ok {
			sop.handleSummary(sum, topics...)
		}
	case MsgError:
		if err, ok := payload.(error); ok {
			sop.handleError(err, topics...)
		}
	}
}

func (sop *stdOutPresenter) handleBegin(topics ...string) {
	if sop.flushStartedTopics() {
		fmt.Println()
	}
	sop.startedTopics = topics
}

func (sop *stdOutPresenter) handleEnd(topics ...string) {
	if sop.flushStartedTopics() {
		fmt.Println()
	}
	topic := path.Join(topics...)

	if _, ok := sop.topicTotals[topic]; ok {
		delete(sop.topicTotals, topic)
	}
	if _, ok := sop.topicPercents[topic]; ok {
		delete(sop.topicPercents, topic)
	}
}

func (sop *stdOutPresenter) handleTotal(total uint64, topics ...string) {
	sop.flushStartedTopics()
	topic := path.Join(topics...)
	sop.topicTotals[topic] = total
}

func (sop *stdOutPresenter) handleProgress(upd uint64, topics ...string) {
	topic := path.Join(topics...)
	if total, ok := sop.topicTotals[topic]; ok {
		pct := float64(upd*100) / float64(total)
		pctStr := fmt.Sprintf("%3.0f", pct)
		if sop.topicPercents[topic] != pctStr {
			fmt.Print("\r")
			presentTopics(topics...)
			fmt.Printf(" %s%%", pctStr)
			if pctStr == "100%" {
				fmt.Println()
			}
			sop.topicPercents[topic] = pctStr
		}
	}
}

func (sop *stdOutPresenter) handleResult(res string, topics ...string) {
	sop.flushStartedTopics()

	fmt.Print("\r")
	presentTopics(topics...)
	fmt.Printf(" %-4s\n", res)
}

func (sop *stdOutPresenter) handleSummary(sum map[string][]string, topics ...string) {
	if sop.flushStartedTopics() {
		fmt.Println()
	}
	if len(topics) == 0 {
		return
	}
	fmt.Println()
	for section, lines := range sum {
		fmt.Println(section)
		for _, line := range lines {
			fmt.Printf(" %s\n", line)
		}
	}
}

func (sop *stdOutPresenter) handleError(err error, topics ...string) {
	sop.flushStartedTopics()
	fmt.Println(err)
}

func (sop *stdOutPresenter) flushStartedTopics() bool {
	needsFlush := len(sop.startedTopics) > 0
	if needsFlush {
		presentTopics(sop.startedTopics...)
		sop.startedTopics = nil
	}
	return needsFlush
}

func presentTopics(topics ...string) {
	if len(topics) == 0 {
		return
	}
	offset := strings.Repeat(" ", len(topics)-1)
	fmt.Print(offset, topics[len(topics)-1])
}
