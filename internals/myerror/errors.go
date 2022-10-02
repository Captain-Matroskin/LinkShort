package myerror

type MyErrors struct {
	Text string
}

func (e *MyErrors) Error() string {
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
	MCreateDBNotConnect           = "db not connect"
	MCreateDBTransactionNotCreate = "transaction setup not create"
	MCreateDBFileNotFound         = "createtables.sql not found"
	MCreateDBFileNotCreate        = "table not create"
	MCreateDBNotCommit            = "transaction setup not commit"
)
