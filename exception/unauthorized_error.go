package exception

type UnAuthoirzedError struct{
	Message string
}

func (e UnAuthoirzedError) Error() string{
	return e.Message
}

func NewUnAuthorizedError(error string) UnAuthoirzedError{
	return UnAuthoirzedError{Message: error}
}