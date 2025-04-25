package service

import (
	_taskModel "github.com/jakkaphatminthana/go-gin/pkg/task/model"
)

type TaskService interface {
	Listing() ([]*_taskModel.Task, error)
	FindById(id uint64) (*_taskModel.Task, error)
}
