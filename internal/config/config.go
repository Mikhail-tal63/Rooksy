package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


type Config struct {
	host string
	jwtSecret string
	jwtExp int64
	port string
}

func Load() *Config{
	 err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system environment variables")
	}

    return &Config{
		host: os.Getenv("HOST"),
		jwtSecret: os.Getenv("SECRET"),
		jwtExp: intEnv("JWT_EXP",(3600*24*7)),
        port: os.Getenv("PORT"),		
		
	}


}

func intEnv(env string, fallback int64) int64{
	if val, ok := os.LookupEnv(env); ok {
		i ,err := strconv.ParseInt(val,10,14)
		if err != nil {
			return fallback
		}
		return i
		
	}
	return fallback
}