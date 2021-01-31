package error

// My own Error type that will help return my customized Error info
type NotFound struct{ msg string }

func (e *NotFound) Error() string { return e.msg }

//IErrorNotFound ...
type IErrorNotFound interface {
	error
}

// Warp the error info in a object
func NewErrorNotFound(msg string) error {
	return &NotFound{msg}
}
