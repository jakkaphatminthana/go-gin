package service

import (
	"github.com/jakkaphatminthana/go-gin/entities"
	_taskModel "github.com/jakkaphatminthana/go-gin/pkg/task/model"
	_taskRepository "github.com/jakkaphatminthana/go-gin/pkg/task/repository"
	"github.com/jakkaphatminthana/go-gin/utils"
)

type taskServiceImpl struct {
	taskRepository _taskRepository.TaskRepository
	logger         utils.Logger
}

func NewTaskServiceImpl(
	taskRepository _taskRepository.TaskRepository,
	logger utils.Logger,
) TaskService {
	return &taskServiceImpl{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

// implement
func (s *taskServiceImpl) Listing() ([]*_taskModel.Task, error) {
	itemList, err := s.taskRepository.Listing()
	if err != nil {
		s.logger.Error("Error listing tasks: ", err)
		return nil, err
	}

	result := s.toTaskResultResponse(itemList)
	return result, nil
}

// implement
func (s *taskServiceImpl) FindById(id uint64) (*_taskModel.Task, error) {
	item, err := s.taskRepository.FindById(id)
	if err != nil {
		s.logger.Error("Error finding task by ID: ", err)
		return nil, err
	}

	return item.ToTaskModel(), nil
}

func (s *taskServiceImpl) toTaskResultResponse(itemEntityList []*entities.Task) []*_taskModel.Task {
	//Mapper entity to model
	taskModelList := make([]*_taskModel.Task, 0)

	for _, task := range itemEntityList {
		taskModelList = append(taskModelList, task.ToTaskModel())
	}

	return taskModelList
}
