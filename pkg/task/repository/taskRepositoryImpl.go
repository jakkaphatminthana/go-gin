package repository

import (
	"github.com/jakkaphatminthana/go-gin/database"
	"github.com/jakkaphatminthana/go-gin/entities"
	_taskException "github.com/jakkaphatminthana/go-gin/pkg/task/exception"
	"github.com/jakkaphatminthana/go-gin/utils"
)

type taskRepositoryImpl struct {
	db     database.Database
	logger utils.Logger
}

func NewTaskRepositoryImpl(db database.Database, logger utils.Logger) TaskRepository {
	return &taskRepositoryImpl{db, logger}
}

// implement
func (r *taskRepositoryImpl) Listing() ([]*entities.Task, error) {
	taskList := make([]*entities.Task, 0)

	query := r.db.Connect().Model(&entities.Task{})

	if err := query.Find(&taskList).Error; err != nil {
		r.logger.Error("Error listing tasks: ", err)
		return nil, &_taskException.TaskListing{}
	}
	return taskList, nil
}

// implement
func (r *taskRepositoryImpl) FindById(id uint64) (*entities.Task, error) {
	task := new(entities.Task)

	if err := r.db.Connect().First(task, id).Error; err != nil {
		r.logger.Errorf("Failed to find task by ID: %s", err.Error())
		return nil, &_taskException.TaskNotFound{ID: id}
	}

	return task, nil
}
