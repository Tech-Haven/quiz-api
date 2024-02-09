package database

import (
	"context"
	"fmt"
	"log"
	"quiz-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func ConnectDatabase(config *config.Config) *mongo.Client {
	db := connectMongo(config)

	err := db.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func connectMongo(c *config.Config) *mongo.Client {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", c.Mongo.User, c.Mongo.Password, c.Mongo.Host, c.Mongo.Port))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
