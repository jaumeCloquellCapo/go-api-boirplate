package error

// My own Error type that will help return my customized Error info
type AlreadyExist struct{ msg string }

// Error implements error interface.
func (e *AlreadyExist) Error() string { return e.msg }

//IErrorAlreadyExist ...
type IErrorAlreadyExist interface {
	error
}

// Error implements error interface.
func NewErrorAlreadyExist(msg string) error {
	return &AlreadyExist{msg}
}
