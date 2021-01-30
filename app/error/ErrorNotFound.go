package error

type ErrorNotFound struct{ msg string }

func (e *ErrorNotFound) Error() string { return e.msg }

//IErrorNotFound ...
type IErrorNotFound interface {
	error
}

//NewErrorNotFound ...
func NewErrorNotFound(msg string) error {
	return &ErrorNotFound{msg}
}
