package config

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// type envConfig struct {
// 	AppPort            string
// 	DbPort             string
// 	DatabaseURL        string
// 	GoogleClientID     string
// 	GoogleClientSecret string
// 	GoogleRedirectURL  string
// 	JWTSaltKey         string
// 	FEOriginalUrl      string
// }

// func (e *envConfig) LoadConfig() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Panic("Error loading .env file")
// 	}

// 	e.AppPort = loadString("APP_PORT", ":8080")
// 	e.DbPort = loadString("DB_PORT", ":5432")
// 	e.DatabaseURL = loadString("DB_BASE_URL", "postgres://postgres:123456@localhost:5432/tasks?sslmode=disable")
// 	e.GoogleClientID = loadString("GOOGLE_CLIENT_ID", "")
// 	e.GoogleClientSecret = loadString("GOOGLE_CLIENT_SECRET", "")
// 	e.GoogleRedirectURL = loadString("GOOGLE_REDIRECT_URL", "http://localhost:3000/dashboard/callback/google")
// 	e.JWTSaltKey = loadString("JWT_SALT_KEY", "hidden_sauce")
// 	e.FEOriginalUrl = loadString("FE_ORIGINAL_URL", "http://localhost:3000")
// }

// var Config envConfig

// func init() {
// 	Config.LoadConfig()
// }

// func loadString(key string, fallback ...string) string {
// 	val, ok := os.LookupEnv(key)
// 	if !ok {
// 		if len(fallback[0]) > 0 {
// 			return fallback[0]
// 		}
// 		log.Panicf("env key: %s is not loaded", key)
// 	}
// 	return val
// }
