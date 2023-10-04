package errors

type UnauthorizedError struct {
}

func (e *UnauthorizedError) Error() string {
	return "Unauthorized"
}
