package main

import (
	"log"

	"github.com/jakkaphatminthana/go-gin/config"
	"github.com/jakkaphatminthana/go-gin/database"
	"github.com/jakkaphatminthana/go-gin/entities"
)

// func taskMigration(tx *gorm.DB) {
// 	if err := tx.Migrator().CreateTable(&entities.Task{}); err != nil {
// 		log.Fatalf("Error creating Task table: %v", err)
// 	}
// }

// func userMigration(tx *gorm.DB) {
// 	if err := tx.Migrator().CreateTable(&entities.User{}); err != nil {
// 		log.Fatalf("Error creating User table: %v", err)
// 	}
// }

// func providerMigration(tx *gorm.DB) {
// 	if err := tx.Migrator().CreateTable(&entities.Provider{}); err != nil {
// 		log.Fatalf("Error creating Provider table: %v", err)
// 	}
// }

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)

	// tx := db.Connect().Begin()

	// // Create table by entities
	// taskMigration(tx)
	// userMigration(tx)
	// providerMigration(tx)

	// if err := tx.Commit().Error; err != nil {
	// 	tx.Rollback()
	// 	log.Fatalf("Error comming migration: %v", err)
	// }

	err := db.Connect().AutoMigrate(
		&entities.Task{},
		&entities.User{},
		&entities.Provider{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

}
