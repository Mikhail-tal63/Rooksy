package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	JwtSecret string
	JwtExp    int64
	Port      string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system environment variables")
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBSSLMode:  os.Getenv("DB_SSL_model"),
		JwtSecret:  os.Getenv("SECRET"),
		JwtExp:     intEnv("JWT_EXP", 3600*24*7),
		Port:       os.Getenv("PORT"),
	}

}

func intEnv(env string, fallback int64) int64 {
	if val, ok := os.LookupEnv(env); ok {
		i, err := strconv.ParseInt(val, 10, 14)
		if err != nil {
			return fallback
		}
		return i

	}
	return fallback
}
