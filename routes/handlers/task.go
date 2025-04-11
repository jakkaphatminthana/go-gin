package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/db"
)

func SaveTask(ctx *gin.Context) {
	var payload db.PostTaskPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body", "data": nil})
		return
	}

	id, err := db.TaskRepository.SaveTask(payload)
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
