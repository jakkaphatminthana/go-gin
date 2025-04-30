package service

import (
	_taskModel "github.com/jakkaphatminthana/go-gin/pkg/task/model"
)

type TaskService interface {
	Listing() ([]*_taskModel.Task, error)
	FindById(id uint64) (*_taskModel.Task, error)
	Create(createReq *_taskModel.TaskCreateReq) (*_taskModel.Task, error)
	Update(taskId uint64, taskUpdateReq *_taskModel.TaskUpdateReq) (*_taskModel.Task, error)
	Delete(taskId uint64) error
}
