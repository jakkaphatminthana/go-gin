package repository

import (
	"github.com/jakkaphatminthana/go-gin/entities"
)

type TaskRepository interface {
	Listing() ([]*entities.Task, error)
	FindById(id uint64) (*entities.Task, error)
}
