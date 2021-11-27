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
	topicTotals           map[string]uint64
	topicPercents         map[string]string
	prevMessage           MessageType
	opportunisticBeforeLF bool
	existingAfterLF       bool
	opportunisticCR       bool
}

func (sop *stdOutPresenter) Handle(msgType MessageType, payload interface{}, topic string) {

	if shouldBreakBefore(msgType, sop.prevMessage) {
		sop.opportunisticBeforeLF = true
	}

	//to overwrite with whitespace to clean up extra characters
	if shouldRewrite(msgType, sop.prevMessage) {
		sop.opportunisticCR = true
	}

	switch msgType {
	case MsgBegin:
		sop.printf("%s ", topic)
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
		if total, ok := sop.topicTotals[topic]; ok && total > 0 {
			if current, ok := payload.(uint64); ok {
				sop.printCurrent(current, topic)
			}
		}
	case MsgResult:
		if result, ok := payload.(string); ok {
			sop.printf("%s %-4s ", topic, result)
		}
	case MsgSummary:
		if summary, ok := payload.(headingSections); ok {
			sop.printSummary(summary)
		}
	case MsgError:
		if err, ok := payload.(error); ok {
			sop.printf("ERROR: %s ", err)
		}
	}

	if shouldBreakAfter(msgType, sop.prevMessage) {
		if !sop.existingAfterLF {
			fmt.Println()
		}
		sop.existingAfterLF = true
	}

	sop.prevMessage = msgType
}

func (sop *stdOutPresenter) printSummary(summary headingSections) {
	sop.opportunisticBeforeLF = true
	sop.existingAfterLF = false

	if summary.heading != "" {
		sop.printf(summary.heading)
		sop.opportunisticBeforeLF = true
		sop.existingAfterLF = false
	}

	for sectionHeading, items := range summary.sections {
		if sectionHeading != "" {
			sop.printf("%s", sectionHeading)
			sop.opportunisticBeforeLF = true
		}
		for _, item := range items {
			sop.printf(" %s", item)
			sop.opportunisticBeforeLF = true
		}
	}
}

func (sop *stdOutPresenter) printCurrent(current uint64, topic string) {
	pct := float64(current*100) / float64(sop.topicTotals[topic])
	pctStr := fmt.Sprintf("%3.0f", pct)
	if topicPct, ok := sop.topicPercents[topic]; !ok || pctStr != topicPct {
		sop.printf("%s %s%% ", topic, pctStr)
	}
	sop.topicPercents[topic] = pctStr
}

func (sop *stdOutPresenter) printf(format string, a ...interface{}) {
	if sop.opportunisticBeforeLF &&
		!sop.existingAfterLF &&
		format != "" {
		fmt.Println()
		sop.opportunisticBeforeLF = false
	}
	if sop.opportunisticCR && format != "" {
		fmt.Print("\r")
		sop.opportunisticCR = false
	}
	fmt.Print(fmt.Sprintf(format, a...))
	if sop.existingAfterLF && format != "" {
		sop.existingAfterLF = false
	}
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
	case MsgSummary:
		return true
	case MsgCurrent:
		return prevMsg == MsgEnd
	}
	return false
}

func shouldBreakAfter(msg, prevMsg MessageType) bool {
	switch msg {
	case MsgError:
		fallthrough
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
