package error

import "errors"

const (
	Success = "success"
	Error   = "error"
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrSqlError            = errors.New("Query Sql Exception")
	ErrTooManyRequests     = errors.New("Too Many Request")
	ErrInvalidToken        = errors.New("Invalid Token")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrForbidden           = errors.New("Forbidden")
)

var GeneralErrors = []error{
	ErrForbidden,
	ErrInternalServerError,
	ErrInvalidToken,
	ErrSqlError,
	ErrTooManyRequests,
	ErrUnauthorized,
}
