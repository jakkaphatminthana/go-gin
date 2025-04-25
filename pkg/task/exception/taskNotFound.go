package exception

import (
	"fmt"
	"net/http"
)

type TaskNotFound struct {
	ID uint64
}

func (e *TaskNotFound) Error() string {
	return fmt.Sprintf("Task with ID %d not found", e.ID)
}

func (e *TaskNotFound) StatusCode() int {
	return http.StatusNotFound
}
