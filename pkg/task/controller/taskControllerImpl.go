package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/pkg/custom"
	_taskModel "github.com/jakkaphatminthana/go-gin/pkg/task/model"
	_taskService "github.com/jakkaphatminthana/go-gin/pkg/task/service"
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

	// validate and convert
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

// implement
func (c *taskControllerImpl) Create(pctx *gin.Context) {
	createReq := new(_taskModel.TaskCreateReq)

	// validate
	customRequest := custom.NewCustomRequest(pctx)
	if err := customRequest.BindBody(createReq); err != nil {
		custom.Error(pctx, err, http.StatusBadRequest)
		return
	}

	// creating
	task, err := c.taskService.Create(createReq)
	if err != nil {
		custom.Error(pctx, err)
		return
	}

	custom.Success(pctx, http.StatusCreated, task)
}

// implement
func (c *taskControllerImpl) Update(pctx *gin.Context) {
	var param = struct {
		ID uint64 `uri:"id" binding:"required"`
	}{}

	// validate and convert (param)
	if err := custom.NewCustomRequest(pctx).BindUri(&param); err != nil {
		custom.Error(pctx, err, http.StatusBadRequest)
		return
	}

	taskUpdateReq := new(_taskModel.TaskUpdateReq)

	// validate and convert (body)
	customRequest := custom.NewCustomRequest(pctx)
	if err := customRequest.BindBody(taskUpdateReq); err != nil {
		custom.Error(pctx, err, http.StatusBadRequest)
		return
	}

	//updating
	task, err := c.taskService.Update(param.ID, taskUpdateReq)
	if err != nil {
		custom.Error(pctx, err)
		return
	}

	custom.Success(pctx, http.StatusCreated, task)
}

// implement
func (c *taskControllerImpl) Delete(pctx *gin.Context) {
	var param = struct {
		ID uint64 `uri:"id" binding:"required"`
	}{}

	// validate and convert (param)
	if err := custom.NewCustomRequest(pctx).BindUri(&param); err != nil {
		custom.Error(pctx, err, http.StatusBadRequest)
		return
	}

	// deleting
	if err := c.taskService.Delete(param.ID); err != nil {
		custom.Error(pctx, err)
		return
	}

	custom.MetaSuccess(pctx, http.StatusOK, "Delete task successful", nil)
}
