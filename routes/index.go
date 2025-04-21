package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/config"
	"github.com/jakkaphatminthana/go-gin/middleware"
	"github.com/jakkaphatminthana/go-gin/routes/handlers"
)

func MounteRoutes() *gin.Engine {
	handler := gin.Default()

	//prevent cros
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Config.FEOriginalUrl},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	taskRoutes := handler.Group("/task", middleware.AuthorizationMiddleware())
	{
		taskRoutes.POST("/", handlers.CreateTask)
		taskRoutes.GET("/", handlers.Tasks)
		taskRoutes.GET("/:id", handlers.Task)
		taskRoutes.PUT("/:id", handlers.UpdateTask)
		taskRoutes.DELETE("/:id", handlers.DeleteTask)
	}

	userLoginRoutes := handler.Group("/login")
	{
		userLoginRoutes.GET("/google", handlers.HandlerGoogleLogin)
	}

	callbackLoginRoutes := handler.Group("/callback")
	{
		callbackLoginRoutes.GET("/google", handlers.HandlerGoogleCallback)
	}

	// handler no route 404
	handler.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})

	return handler
}
