package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/config"
	"github.com/jakkaphatminthana/go-gin/db"
)

func main() {
	db.InitDB()
	handler := gin.Default()
	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: handler,
	}
	// close the database connection when the server is closed
	defer db.DB.Close(context.Background())

	// start the server and listen
	server.ListenAndServe()
}
