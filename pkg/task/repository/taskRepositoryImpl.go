package repository

import (
	"github.com/jakkaphatminthana/go-gin/database"
	"github.com/jakkaphatminthana/go-gin/entities"
	"github.com/jakkaphatminthana/go-gin/pkg/custom"
	_taskException "github.com/jakkaphatminthana/go-gin/pkg/task/exception"
	_taskModel "github.com/jakkaphatminthana/go-gin/pkg/task/model"
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

// implement
func (r *taskRepositoryImpl) Create(taskEntity *entities.Task) (*entities.Task, error) {
	data := new(entities.Task)

	if err := r.db.Connect().Create(taskEntity).Scan(data).Error; err != nil {
		r.logger.Errorf("Creating task failed: %s", err.Error())
		return nil, custom.ErrorInternalServerError("create task failed")
	}

	return data, nil
}

// implement
func (r *taskRepositoryImpl) Update(taskId uint64, taskUpdateReq *_taskModel.TaskUpdateReq) (uint64, error) {
	if err := r.db.Connect().Model(&entities.Task{}).Where("id = ?", taskId).Updates(taskUpdateReq).Error; err != nil {
		r.logger.Errorf("Updating task failed: %s", err.Error())
		return 0, &_taskException.TaskUpdateFailed{ID: taskId}
	}
	return taskId, nil
}

// implement
func (r *taskRepositoryImpl) Delete(taskId uint64) error {
	if err := r.db.Connect().Where("id = ?", taskId).Delete(&entities.Task{}).Error; err != nil {
		r.logger.Errorf("Deleting task filed: %s", err.Error())
		return &_taskException.TaskDeleteFailed{ID: taskId}
	}
	return nil
}
