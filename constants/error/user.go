package error

import "errors"

var (
	ErrUserNotFound         = errors.New("User Not Found")
	ErrPasswordIncorrect    = errors.New("Password Incorrect")
	ErrUsernameExist        = errors.New("Username Exist")
	ErrPasswordDoesNotMatch = errors.New("Password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordDoesNotMatch,
	ErrPasswordIncorrect,
	ErrUsernameExist,
}
