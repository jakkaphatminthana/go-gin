package model

import (
	"time"

	"github.com/jakkaphatminthana/go-gin/types"
)

type (
	Task struct {
		ID        uint64    `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
	}

	TaskCreateReq struct {
		Title   string           `json:"title" validate:"required,max=100"`
		Content string           `json:"content" validate:"omitempty,max=255"`
		Status  types.TaskStatus `json:"status" validate:"required,oneof='todo' 'doing' 'done'"`
	}

	TaskUpdateReq struct {
		Title   string           `json:"title" validate:"required,max=100"`
		Content string           `json:"content" validate:"omitempty,max=255"`
		Status  types.TaskStatus `json:"status" validate:"required,oneof='todo' 'doing' 'done'"`
	}
)
