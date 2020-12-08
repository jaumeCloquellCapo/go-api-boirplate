package error

type ErrorNotFound struct{ msg string }

func (e *ErrorNotFound) IsNotFound() bool { return true }
func (e *ErrorNotFound) Error() string    { return e.msg }

type IErrorNotFound interface {
	error
	IsNotFound() bool
}
func NewErrorNotFound(msg string) error {
	return &ErrorNotFound{msg}
}
