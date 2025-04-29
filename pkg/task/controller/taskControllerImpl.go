package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/pkg/custom"
	_taskService "github.com/jakkaphatminthana/go-gin/pkg/task/service"
	"github.com/jakkaphatminthana/go-gin/utils"
)

type taskControllerImpl struct {
	taskService _taskService.TaskService
}

func NewTaskControllerImpl(
	taskService _taskService.TaskService,
) TaskController {
	return &taskControllerImpl{
		taskService: taskService,
	}
}

// implement
func (c *taskControllerImpl) Listing(pctx *gin.Context) {
	taskModelList, err := c.taskService.Listing()
	if err != nil {
		custom.Error(pctx, err, http.StatusInternalServerError)
		return
	}

	custom.Success(pctx, http.StatusOK, taskModelList)
}

// implement
func (c *taskControllerImpl) FindById(pctx *gin.Context) {
	var param = struct {
		ID uint64 `uri:"id" binding:"required"`
	}{}

	if err := custom.NewCustomRequest(pctx).BindUri(&param); err != nil {
		custom.Error(pctx, err, http.StatusBadRequest)
		return
	}

	taskModel, err := c.taskService.FindById(param.ID)
	if err != nil {
		custom.Error(pctx, err)
		return
	}

	custom.Success(pctx, http.StatusOK, taskModel)
}
