package error

// My own Error type that will help return my customized Error info
type NotFound struct{ msg string }

// Error implements error interface.
func (e *NotFound) Error() string { return e.msg }

//IErrorNotFound ...
type IErrorNotFound interface {
	error
}

// Error implements error interface.
func NewErrorNotFound(msg string) error {
	return &NotFound{msg}
}
