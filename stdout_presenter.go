package utr

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
	case MsgBegin:
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
	case MsgSuccess:
		if success, ok := payload.(bool); ok {
			sop.handleSuccess(success, topics...)
		}
	case MsgSummary:
		if sum, ok := payload.(map[string][]string); ok {
			sop.handleSummary(sum, topics...)
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
	newLine := false
	if _, ok := sop.topicTotals[topic]; ok {
		delete(sop.topicTotals, topic)
		newLine = true
	}
	if _, ok := sop.topicPercents[topic]; ok {
		delete(sop.topicPercents, topic)
		newLine = true
	}
	if newLine {
		fmt.Println()
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
			sop.topicPercents[topic] = pctStr
		}
	}
}

func (sop *stdOutPresenter) handleSuccess(success bool, topics ...string) {
	fmt.Print("\r")
	presentTopics(topics...)
	str := "OK"
	if !success {
		str = "FAIL"
	}
	fmt.Printf(" %-4s", str)
}

func (sop *stdOutPresenter) handleSummary(sum map[string][]string, topics ...string) {
	if sop.flushStartedTopics() {
		fmt.Println()
	}
	if len(topics) > 0 {
		return
	}
	fmt.Println()
	for section, lines := range sum {
		fmt.Println(section)
		for _, line := range lines {
			fmt.Println(line)
		}
	}
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
