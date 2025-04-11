package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/db"
	"net/http"
)

func SaveTask(ctx *gin.Context) {
	var payload db.PostTaskPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body", "data": nil})
		return
	}

	id, err := db.TaskRepository.SaveTaskQuery(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save task", "data": nil})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "data": gin.H{"id": id}})
}
