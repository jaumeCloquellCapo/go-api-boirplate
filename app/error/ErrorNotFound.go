package error

type ErrorNotFound struct{ msg string }

func (e *ErrorNotFound) Error() string { return e.msg }

type IErrorNotFound interface {
	error
}

func NewErrorNotFound(msg string) error {
	return &ErrorNotFound{msg}
}
