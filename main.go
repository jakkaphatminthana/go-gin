package main

import (
	"github.com/jakkaphatminthana/go-gin/config"
	"github.com/jakkaphatminthana/go-gin/database"
	_ "github.com/jakkaphatminthana/go-gin/docs"
	"github.com/jakkaphatminthana/go-gin/server"
)

//	@title			GoGin API
//	@version		1.0
//	@description	This is a sample server.
//	@termsOfService	http://swagger.io/terms/

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/v1
func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)
	server := server.NewGinServer(conf, db)

	server.Start()
}
