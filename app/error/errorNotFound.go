package error

// My own Error type that will help return my customized Error info
type ErrorNotFound struct{ msg string }

func (e *ErrorNotFound) Error() string { return e.msg }

//IErrorNotFound ...
type IErrorNotFound interface {
	error
}

// Warp the error info in a object
func NewErrorNotFound(msg string) error {
	return &ErrorNotFound{msg}
}
