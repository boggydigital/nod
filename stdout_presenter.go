package nod

import "fmt"

func EnableStdOutPresenter() {
	sop := &stdOutPresenter{
		topicTotals:   make(map[string]uint64),
		topicPercents: make(map[string]string),
		prevMessage:   MsgNone,
	}
	HandleFunc(sop, StdOutTypes()...)
}

type stdOutPresenter struct {
	topicTotals     map[string]uint64
	topicPercents   map[string]string
	prevMessage     MessageType
	opportunisticLF bool
	opportunisticCR bool
}

func (sop *stdOutPresenter) Handle(msgType MessageType, payload interface{}, topic string) {

	if shouldBreakBefore(msgType, sop.prevMessage) {
		sop.opportunisticLF = true
	}

	if shouldRewrite(msgType, sop.prevMessage) {
		sop.opportunisticCR = true
	}

	switch msgType {
	case MsgBegin:
		sop.printf(topic + " ")
	case MsgEnd:
		if _, ok := sop.topicTotals[topic]; ok {
			delete(sop.topicTotals, topic)
		}
		if _, ok := sop.topicPercents[topic]; ok {
			delete(sop.topicPercents, topic)
		}
	case MsgTotal:
		if total, ok := payload.(uint64); ok {
			sop.topicTotals[topic] = total
		}
	case MsgCurrent:
		if current, ok := payload.(uint64); ok {
			pct := float64(current*100) / float64(sop.topicTotals[topic])
			pctStr := fmt.Sprintf("%3.0f", pct)
			if topicPct, ok := sop.topicPercents[topic]; !ok || pctStr != topicPct {
				sop.printf("%s %s%% ", topic, pctStr)
			}
			sop.topicPercents[topic] = pctStr
		}
	case MsgResult:
		if result, ok := payload.(string); ok {
			sop.printf("%s %-4s ", topic, result)
		}
	case MsgError:
		if err, ok := payload.(error); ok {
			sop.printf("%-4s ", err)
		}
	}

	if shouldBreakAfter(msgType, sop.prevMessage) {
		sop.opportunisticLF = true
	}

	sop.prevMessage = msgType
}

func (sop *stdOutPresenter) printf(format string, a ...interface{}) {
	if sop.opportunisticLF {
		fmt.Println()
		sop.opportunisticLF = false
	}
	if sop.opportunisticCR {
		fmt.Print("\r")
		sop.opportunisticCR = false
	}
	fmt.Print(fmt.Sprintf(format, a...))
}

func shouldBreakBefore(msg, prevMsg MessageType) bool {
	switch msg {
	case MsgBegin:
		switch prevMsg {
		case MsgTotal:
			fallthrough
		case MsgBegin:
			fallthrough
		case MsgCurrent:
			return true
		}
	}
	return false
}

func shouldBreakAfter(msg, prevMsg MessageType) bool {
	switch msg {
	case MsgEnd:
		return true
	}
	return false
}

func shouldRewrite(msg, prevMsg MessageType) bool {
	switch msg {
	case MsgResult:
		fallthrough
	case MsgCurrent:
		return true
	}
	return false
}
