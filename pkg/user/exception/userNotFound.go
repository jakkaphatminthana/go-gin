package exception

import (
	"fmt"
	"net/http"
)

type UserNotFound struct {
	ID uint64
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("user ID: %d not found", e.ID)
}

func (e *UserNotFound) StatusCode() int {
	return http.StatusNotFound
}
