package model

import "time"

type (
	Task struct {
		ID        uint64    `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
	}
)
