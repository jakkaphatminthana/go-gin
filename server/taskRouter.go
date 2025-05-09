package server

import (
	"github.com/jakkaphatminthana/go-gin/middleware"
	_taskController "github.com/jakkaphatminthana/go-gin/pkg/task/controller"
	_taskRepository "github.com/jakkaphatminthana/go-gin/pkg/task/repository"
	_taskService "github.com/jakkaphatminthana/go-gin/pkg/task/service"
	"github.com/jakkaphatminthana/go-gin/utils"
)

func (s *ginServer) initTaskRouter(m *middleware.AuthorizationMiddleware) {

	taskRepository := _taskRepository.NewTaskRepositoryImpl(s.db, utils.GetLogger())
	taskService := _taskService.NewTaskServiceImpl(taskRepository, utils.GetLogger())
	taskController := _taskController.NewTaskControllerImpl(taskService)

	//public endpoint
	publicRouter := s.engine.Group("/v1/task")
	publicRouter.GET("/", taskController.Listing)
	publicRouter.GET("/:id", taskController.FindById)

	//private endpoint
	privateRouter := s.engine.Group("/v1/task")
	privateRouter.Use(m.Handler())

	privateRouter.POST("/", taskController.Create)
	privateRouter.PATCH("/:id", taskController.Update)
	privateRouter.DELETE("/:id", taskController.Delete)
}
