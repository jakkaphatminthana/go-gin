package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jakkaphatminthana/go-gin/config"
)

var DB *pgx.Conn

func InitDB() {
	// url for connect to database postgres
	urlExample := config.Config.DatabaseURL
	var err error

	// connect to database
	DB, err = pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// check the connection
	err = DB.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database")
}
