package config

import (
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server   `mapstructure:"server" validate:"required"`
		OAuth2   *OAuth2   `mapstructure:"oauth2" validate:"required"`
		Database *Database `mapstructure:"database" validate:"required"`
		EVNValue *EVNValue
	}
	Server struct {
		Port         string        `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
	}
	OAuth2 struct {
		GoogleClientID     string   `mapstructure:"googleClientID" validate:"required"`
		GoogleClientSecret string   `mapstructure:"googleClientSecret" validate:"required"`
		GoogleRedirectURL  string   `mapstructure:"googleRedirectURL" validate:"required"`
		Scopes             []string `mapstructure:"scopes" validate:"required"`
		UserInfoUrl        string   `mapstructure:"userInfoUrl" validate:"required"`
	}
	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbName" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}

	EVNValue struct {
		JWTSaltKey    string
		FEOriginalUrl string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func replaceEnvVariables(value string) string {
	re := regexp.MustCompile(`\${(.*?)}`)
	return re.ReplaceAllStringFunc(value, func(match string) string {
		key := strings.Trim(match, "${}")
		if val, exists := os.LookupEnv(key); exists {
			return val
		}
		return match
	})
}

func ConfigGetting() *Config {
	once.Do(func() {
		// load env
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found, using system environment variables")
		}

		// Initialize Viper
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./bin")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		// server.port -> SERVER_PORT (environment variable)
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		// Read config.yaml
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		// Replace environment variables in the config
		for _, key := range viper.AllKeys() {
			if strValue, ok := viper.Get(key).(string); ok {
				viper.Set(key, replaceEnvVariables(strValue))
			}
		}

		// Watch for config changes
		viper.WatchConfig()

		// convert config to struct
		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		validating := validator.New()
		// validating by struct (validate:"required")
		if err := validating.Struct(configInstance); err != nil {
			panic(err)
		}

		// Load config environment variables
		configInstance.EVNValue = LoadEnvValues()
	})
	return configInstance
}

func LoadEnvValues() *EVNValue {
	jwtKey := os.Getenv("JWT_SALT_KEY")
	feUrl := os.Getenv("FE_ORIGINAL_URL")

	if jwtKey == "" || feUrl == "" {
		log.Fatal("Missing required environment variables")
	}

	return &EVNValue{
		JWTSaltKey:    jwtKey,
		FEOriginalUrl: feUrl,
	}
}
