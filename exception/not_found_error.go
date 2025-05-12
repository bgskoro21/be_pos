package exception

type NotFoundError struct{
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(error string) NotFoundError{
	return NotFoundError{Message: error}
}