package exceptions

type UnauthorizedError struct {
}

func (e *UnauthorizedError) Error() string {
	return "Unauthorized"
}

func (e *UnauthorizedError) Code() int {
	return 401
}
