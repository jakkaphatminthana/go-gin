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

	taskRoutes := handler.Group("/task")
	{
		taskRoutes.POST("/", handlers.SaveTask)
		taskRoutes.GET("/", handlers.Tasks)
	}

	// handler no route 404
	handler.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})

	return handler
}
