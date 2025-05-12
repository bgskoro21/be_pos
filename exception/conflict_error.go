package exception

type ConflictError struct{
	Message string
}

func (e ConflictError) Error() string {
	return e.Message
}

func NewConflictError(error string) ConflictError{
	return ConflictError{Message: error}
}