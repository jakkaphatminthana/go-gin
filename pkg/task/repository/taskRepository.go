package repository

import (
	"github.com/jakkaphatminthana/go-gin/entities"
	_taskModel "github.com/jakkaphatminthana/go-gin/pkg/task/model"
)

type TaskRepository interface {
	Listing() ([]*entities.Task, error)
	FindById(id uint64) (*entities.Task, error)
	Create(taskEntity *entities.Task) (*entities.Task, error)
	Update(taskId uint64, taskUpdateReq *_taskModel.TaskUpdateReq) (uint64, error)
	Delete(taskId uint64) error
}
