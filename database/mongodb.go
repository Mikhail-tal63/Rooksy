package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(url, name string) (*mongo.Database, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(cxt, clientOptions)
	if err != nil {

		return nil, err
	}
	err = client.Ping(cxt, nil)
	if err != nil {
		return nil, err
	}
	log.Println("connected to mongoDB")
	return client.Database(name), nil
}
