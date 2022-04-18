package exception

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func GetErrorInfo(s string) error {
	return &errorString{s}
}
