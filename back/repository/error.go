package repository

type IDError struct {
	Msg string
}

func (e *IDError) Error() string {
	return "IDError: " + e.Msg
}

type IOError struct {
	Msg string
}

func (e *IOError) Error() string {
	return "IOError: " + e.Msg
}

type NotFoundRecordError struct {
	Msg string
}

func (e *NotFoundRecordError) Error() string {
	return "NotFoundRecordError: " + e.Msg
}
