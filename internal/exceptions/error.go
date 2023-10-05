package exceptions

type Error interface {
	Error() string
	Code() int
}

type UnknownError struct {
}

func (e *UnknownError) Error() string {
	return "Unknown error"
}

func (e *UnknownError) Code() int {
	return 500
}
