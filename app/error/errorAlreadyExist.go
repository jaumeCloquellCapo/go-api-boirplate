package error

// My own Error type that will help return my customized Error info
type AlreadyExist struct{ msg string }

func (e *AlreadyExist) Error() string { return e.msg }

//IErrorNotFound ...
type IErrorAlreadyExist interface {
	error
}

// Warp the error info in a object
func NewErrorAlreadyExist(msg string) error {
	return &AlreadyExist{msg}
}
