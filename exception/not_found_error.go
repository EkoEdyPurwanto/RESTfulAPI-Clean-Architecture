package exception

type NotFoundError struct {
	ErrorMsg string
}

func (e NotFoundError) Error() string {
	return e.ErrorMsg
}

func NewNotFoundError(errorMsg string) NotFoundError {
	return NotFoundError{ErrorMsg: errorMsg}
}
