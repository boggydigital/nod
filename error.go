package nod

func Error(err error) error {
	dispatch(MsgError, nil, err.Error())
	return err
}
