package nod

func Error(err error) error {
	dispatch(MsgError, nil, err.Error())
	return err
}

func ErrorStr(s string) string {
	dispatch(MsgError, nil, s)
	return s
}