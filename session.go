package nod

type session struct {
	active bool
}

func SessionBegin() ActCloser {
	dispatch(MsgSessionBegin, nil, "")
	return &session{active: true}
}

func (s *session) End() {
	if s.active {
		dispatch(MsgSessionEnd, nil, "")
		s.active = false
	}
}

func (s *session) EndWithResult(result string) {
	if s.active {
		dispatch(MsgResult, result, "")
	}
}

func (s *session) EndWithError(err error) error {
	if s.active {
		dispatch(MsgError, err, "")
	}
	return err
}

func (s *session) EndWithSummary(summary map[string][]string) {
	if s.active {
		dispatch(MsgSummary, summary, "")
	}
}
