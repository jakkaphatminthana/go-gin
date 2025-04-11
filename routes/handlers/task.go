package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/db"
)

func CreateTask(ctx *gin.Context) {
	var payload db.CreateTaskReq
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body", "data": nil})
		return
	}

	id, err := db.TaskRepository.CreateTask(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save task", "data": nil})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"error": nil, "data": gin.H{"id": id}})
}

func Tasks(ctx *gin.Context) {
	tasks, err := db.TaskRepository.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch tasks", "data": nil})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": nil, "data": tasks})
}

func Task(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required", "data": nil})
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a valid integer", "data": nil})
		return
	}

	task, err := db.TaskRepository.GetTaskByID(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch task", "data": nil})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": nil, "data": task})
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required", "data": nil})
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a valid integer", "data": nil})
		return
	}

	var payload db.UpdateTaskReq
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body", "data": nil})
		return
	}

	task, err := db.TaskRepository.GetTaskByID(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": nil})
		return
	}

	if payload.Title != "" {
		task.Title = payload.Title
	}
	if payload.Content != "" {
		task.Content = payload.Content
	}
	if payload.Status != "" {
		task.Status = payload.Status
	}

	if updateDataErr := db.TaskRepository.UpdateTask(intID, payload); updateDataErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update task", "data": nil})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": nil, "data": payload})
}
