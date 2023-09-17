package status

var _ error = &Error{}

type Error struct {
	ogErr      error
	StatusCode int
}

func (e *Error) Error() string {
	return e.ogErr.Error()
}
