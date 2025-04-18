package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	AppPort      string
	DbPort       string
	DatabaseURL  string
	ClientID     string
	ClientSecret string
}

func (e *envConfig) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	e.AppPort = loadString("APP_PORT", ":8080")
	e.DbPort = loadString("DB_PORT", ":5432")
	e.DatabaseURL = loadString("DB_BASE_URL", "postgres://postgres:123456@localhost:5432/tasks?sslmode=disable")
	e.ClientID = loadString("CLIENT_ID")
	e.ClientSecret = loadString("CLIENT_SECRET")
}

var Config envConfig

func init() {
	Config.LoadConfig()
}

func loadString(key string, fallback ...string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		if len(fallback[0]) > 0 {
			return fallback[0]
		}
		log.Panicf("%s is not loaded", key)
	}
	return val
}
