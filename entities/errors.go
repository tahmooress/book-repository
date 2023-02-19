package entities

type Type string

const (
	Internal       Type = "something went wrong, try again!"
	UnAuthorize    Type = "authorization failed"
	DuplicateUser  Type = "user with this email already exist"
	InvalidPssword Type = "password is not valid"
	NotFound       Type = "requested data is not exist"
)

type Err struct {
	Type   Type
	Reason error
}

func (e *Err) Error() string {
	return e.Reason.Error()
}

func (e *Err) Message() string {
	return string(e.Type)
}

func NewError(err error, message Type) *Err {
	return &Err{
		Type:   message,
		Reason: err,
	}
}
