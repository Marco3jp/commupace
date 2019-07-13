package module

type UUIDCreateError struct {
	Msg string
}

func (e *UUIDCreateError) Error() string {
	return "UUIDCreateError: " + e.Msg
}

type ExpiredRefreshTokenError struct {
	Msg string
}

func (e *ExpiredRefreshTokenError) Error() string {
	return "ExpiredRefreshTokenError: " + e.Msg
}

type InvalidRequestError struct {
	Msg string
}

func (e *InvalidRequestError) Error() string {
	return "InvalidRequestError: " + e.Msg
}
