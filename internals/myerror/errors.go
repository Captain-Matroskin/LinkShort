package myerror

type Errors struct {
	Text string
}

func (e *Errors) Error() string {
	return e.Text
}

type CheckError struct {
	RequestId int
}

// Error of server
const (
	ErrAtoi            = "func Atoi convert string in int"
	IntNil             = 0
	ErrNotStringAndInt = "expected type string or int"
)

// Error of main
const (
	MCreateDBNotConnect = "db not connect"
)
