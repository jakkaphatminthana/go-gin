package db

import (
	"context"
	"time"
)

type TaskRepositoryS struct{}

var TaskRepository = TaskRepositoryS{}

type CreateTaskReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  string `json:"status"`
}

func (t TaskRepositoryS) CreateTask(payload CreateTaskReq) (int, error) {
	if payload.Status == "" {
		payload.Status = "todo"
	}

	var id int
	query := `INSERT INTO tasks (title, content, status) VALUES ($1, $2, $3) RETURNING id;`

	if err := DB.QueryRow(context.Background(), query, payload.Title, payload.Content, payload.Status).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (t TaskRepositoryS) GetTasks() ([]Task, error) {
	var tasks []Task

	query := `SELECT id, title, content, status, created_at FROM tasks ORDER BY created_at DESC LIMIT 10;`

	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item Task
		if err := rows.Scan(&item.ID, &item.Title, &item.Content, &item.Status, &item.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, item)
	}

	return tasks, nil
}

func (t TaskRepositoryS) GetTaskByID(id int) (Task, error) {
	var task Task

	query := `SELECT id, title, content, status, created_at FROM tasks WHERE id = $1;`
	if err := DB.QueryRow(context.Background(), query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Content,
		&task.Status,
		&task.CreatedAt,
	); err != nil {
		return Task{}, err
	}

	return task, nil
}

type UpdateTaskReq struct {
	Title   string `json:"title" binding:"max=100"`
	Content string `json:"content" binding:"max=255"`
	Status  string `json:"status"`
}

func (t TaskRepositoryS) UpdateTask(id int, payload UpdateTaskReq) error {
	query := `UPDATE tasks SET title = $1, content = $2, status = $3 WHERE id = $4;`
	if _, err := DB.Exec(context.Background(), query, payload.Title, payload.Content, payload.Status, id); err != nil {
		return err
	}

	return nil
}
