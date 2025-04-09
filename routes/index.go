package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/routes/handlers"
)

func MounteRoutes() *gin.Engine {
	handler := gin.Default()

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	handler.POST("/task", handlers.SaveTask)

	return handler
}
