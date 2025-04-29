package custom

import "net/http"

type AppError interface {
	error
	StatusCode() int
}

type appErrorImpl struct {
	statusCode int
	message    string
}

func (e *appErrorImpl) Error() string {
	return e.message
}

func (e *appErrorImpl) StatusCode() int {
	return e.statusCode
}

func newError(statusCode int, message string) AppError {
	return &appErrorImpl{
		statusCode: statusCode,
		message:    message,
	}
}

// shortcut
func ErrorBadRequest(message string) AppError {
	return newError(http.StatusBadRequest, message)
}

func ErrorNotFound(message string) AppError {
	return newError(http.StatusNotFound, message)
}

func ErrorUnauthorized(message string) AppError {
	return newError(http.StatusUnauthorized, message)
}

func ErrorInternalServerError(message string) AppError {
	return newError(http.StatusInternalServerError, message)
}
