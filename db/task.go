package db

import (
	"context"
)

type Task struct{}

var TaskRepository = Task{}

type PostTaskPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}

func (t Task) SaveTaskQuery(payload PostTaskPayload) (int, error) {
	if payload.Status == "" {
		payload.Status = "todo"
	}

	var id int
	query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id;`

	if err := DB.QueryRow(context.Background(), query, payload.Title, payload.Description, payload.Status).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}
