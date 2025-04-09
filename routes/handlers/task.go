package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostTaskReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}

func SaveTask(ctx *gin.Context) {
	var payload PostTaskReq

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body", "data": nil})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": nil, "data": payload})
}
