package exception

import (
	"fmt"
	"net/http"
)

type TaskUpdateFailed struct {
	ID uint64
}

func (e *TaskUpdateFailed) Error() string {
	return fmt.Sprintf("updating task id: %d failed", e.ID)
}

func (e *TaskUpdateFailed) StatusCode() int {
	return http.StatusInternalServerError
}
