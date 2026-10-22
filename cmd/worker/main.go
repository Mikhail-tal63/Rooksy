package main

import (
	"log"

	"github.com/Mikhail-Tal63/Rooksy/cmd/api"
	"github.com/Mikhail-Tal63/Rooksy/database"
	"github.com/Mikhail-Tal63/Rooksy/internal/config"
)

func main() {
	cfg := config.Load()

	db, err := database.ConnectMongoDB(cfg.Host, cfg.DBName)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	server := api.NewAPIServer(db, cfg.Port)
	if err := server.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
