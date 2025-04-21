package main

import (
	"github.com/jakkaphatminthana/go-gin/config"
	"github.com/jakkaphatminthana/go-gin/database"
	"github.com/jakkaphatminthana/go-gin/server"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)
	server := server.NewGinServer(conf, db)

	server.Start()
}
