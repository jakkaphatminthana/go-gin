package exception

import (
	"fmt"
	"net/http"
)

type TaskDeleteFailed struct {
	ID uint64
}

func (e *TaskDeleteFailed) Error() string {
	return fmt.Sprintf("delete task id: %d failed", e.ID)
}

func (e *TaskDeleteFailed) StatusCode() int {
	return http.StatusInternalServerError
}
