package main

import (
	"context"
	"github.com/jakkaphatminthana/go-gin/config"
	"github.com/jakkaphatminthana/go-gin/db"
	"github.com/jakkaphatminthana/go-gin/routes"
	"net/http"
)

func main() {
	db.InitDB()
	handler := routes.MounteRoutes()

	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: handler,
	}
	// close the database connection when the server is closed
	defer db.DB.Close(context.Background())

	// start the server and listen
	server.ListenAndServe()
}
