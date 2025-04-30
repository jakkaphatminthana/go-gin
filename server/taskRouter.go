package server

import (
	_taskController "github.com/jakkaphatminthana/go-gin/pkg/task/controller"
	_taskRepository "github.com/jakkaphatminthana/go-gin/pkg/task/repository"
	_taskService "github.com/jakkaphatminthana/go-gin/pkg/task/service"
	"github.com/jakkaphatminthana/go-gin/utils"
)

func (s *ginServer) initTaskRouter() {
	router := s.engine.Group("/v1/task")

	taskRepository := _taskRepository.NewTaskRepositoryImpl(s.db, utils.GetLogger())
	taskService := _taskService.NewTaskServiceImpl(taskRepository, utils.GetLogger())
	taskController := _taskController.NewTaskControllerImpl(taskService)

	//endpoint
	router.GET("/", taskController.Listing)
	router.GET("/:id", taskController.FindById)
	router.POST("/", taskController.Create)
	router.PATCH("/:id", taskController.Update)
	router.DELETE("/:id", taskController.Delete)
}
