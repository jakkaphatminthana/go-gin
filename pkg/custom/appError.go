package custom

import "net/http"

type AppError interface {
	error
	StatusCode() int
}

type AppErrorImpl struct {
	statusCode int
	message    string
}

func (e *AppErrorImpl) Error() string {
	return e.message
}

func (e *AppErrorImpl) StatusCode() int {
	return e.statusCode
}

func newError(statusCode int, message string) *AppErrorImpl {
	return &AppErrorImpl{
		statusCode: statusCode,
		message:    message,
	}
}

// shortcut
func ErrorBadRequest(message string) *AppErrorImpl {
	return newError(http.StatusBadRequest, message)
}

func ErrorNotFound(message string) *AppErrorImpl {
	return newError(http.StatusNotFound, message)
}

func ErrorUnauthorized(message string) *AppErrorImpl {
	return newError(http.StatusUnauthorized, message)
}

func ErrorInternalServerError(message string) *AppErrorImpl {
	return newError(http.StatusInternalServerError, message)
}
