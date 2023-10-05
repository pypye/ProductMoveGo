package exceptions

type WrongIdentityError struct {
}

func (e *WrongIdentityError) Error() string {
	return "Wrong username or password"
}

func (e *WrongIdentityError) Code() int {
	return 400
}
