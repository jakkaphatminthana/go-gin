package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	AppPort     string
	DbPort      string
	DatabaseURL string
}

func (e *envConfig) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	e.AppPort = loadString("APP_PORT", ":8080")
	e.DbPort = loadString("DB_PORT", ":5432")
	e.DatabaseURL = loadString("DATABASE_URL", "postgres://postgres:123456@localhost:5432/tasks?sslmode=disable")
}

var Config envConfig

func init() {
	Config.LoadConfig()
}

func loadString(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("%s is not loaded", key)
		return fallback
	}
	return val
}
