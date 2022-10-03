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

type ResultError struct {
	Status  int    `json:"status"`
	Explain string `json:"explain,omitempty"`
}

// Error of server
const (
	ErrDB              = "database is not responding"
	ErrAtoi            = "func Atoi convert string in int"
	IntNil             = 0
	ErrNotStringAndInt = "expected type string or int"
	ErrUnmarshal       = "unmarshal json"
	ErrMarshal         = "marshaling in json"
	ErrCheck           = "err check"
	ErrEncode          = "Encode"
)

// Error of main
const (
	MCreateDBNotConnect           = "db not connect"
	MCreateDBTransactionNotCreate = "transaction setup not create"
	MCreateDBFileNotFound         = "createtables.sql not found"
	MCreateDBFileNotCreate        = "table not create"
	MCreateDBNotCommit            = "transaction setup not commit"
)

// Error of LinkShort
const (
	LSHCreateLinkShortTransactionNotCreate = "transaction Create Link Short not create"
	LSHCreateLinkShortNotInsert            = "Link short not insert"
	LSHCreateLinkShortNotCommit            = "Link short not commit"
	LSHCreateLinkShortNotInsertUniqueDB    = "ERROR: duplicate key value violates unique constraint \"link_link_key\" (SQLSTATE 23505)"
	LSHCreateLinkShortNotInsertUnique      = "link is not unique"
)
