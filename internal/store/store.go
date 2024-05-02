package store

import (
	"context"
	"dttg/internal/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Store struct {
	MongoClient *mongo.Client
}

func NewStore(context *context.Context, config config.Config) (*Store, error) {
	var err error
	store := &Store{}
	store.MongoClient, err = initMongo(context, config)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func initMongo(context *context.Context, config config.Config) (*mongo.Client, error) {
	connectionString := fmt.Sprintf("mongodb://%s:%s/%s",
		config.MongoDb.HOST,
		config.MongoDb.PORT,
		config.MongoDb.DB,
	)
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(*context, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(*context, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client, nil
}
