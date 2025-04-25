package exception

import "net/http"

type TaskListing struct{}

func (e *TaskListing) Error() string {
	return "Task listing failed"
}

func (e *TaskListing) StatusCode() int {
	return http.StatusInternalServerError
}
