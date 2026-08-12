package custom

import (
	_taskModel "github.com/jakkaphatminthana/go-gin/pkg/task/model"
)

// Error Response
type (
	DocsErrorBadRequest struct {
		Error struct {
			Status  int               `json:"status" example:"400"`
			Message string            `json:"message" example:"Bad request"`
			Details map[string]string `json:"details,omitempty"`
		} `json:"error"`
	}

	DocsErrorNotFound struct {
		Error struct {
			Status  int               `json:"status" example:"404"`
			Message string            `json:"message" example:"Not Found"`
			Details map[string]string `json:"details,omitempty"`
		} `json:"error"`
	}

	DocsErrorInternalServerError struct {
		Error struct {
			Status  int               `json:"status" example:"500"`
			Message string            `json:"message" example:"Internal Server Error"`
			Details map[string]string `json:"details,omitempty"`
		} `json:"error"`
	}
)

// Success Response
type (
	Gen[T any] struct {
		Data T `json:"data"`
	}
	DocsTasks struct {
		Data []_taskModel.Task `json:"data"`
	}
	DocsTask struct {
		Data _taskModel.Task `json:"data"`
	}
)
