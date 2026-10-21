package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	host string
	jwtSecret string
	jwtExp string
	port int
}

func Load() *Config{
	 err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system environment variables")
	}

    return &Config{
		host: os.Getenv("HOST"),
		jwtSecret: os.Getenv("SECRET"),
		jwtExp: os.Getenv("JWT_EXP"),
		
	}


}