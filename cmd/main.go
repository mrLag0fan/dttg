package main

import (
	"context"
	"dttg/internal/config"
	"dttg/internal/repository/impl"
	impl2 "dttg/internal/service/impl"
	"dttg/internal/store"
	"dttg/pkg/api"
	"dttg/pkg/api/controller"
	"log"
)

func main() {
	myContext := context.Background()
	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Fatalf("Can`t read congif from ENV: %v", err)
	}
	newStore, err := store.NewStore(&myContext, *cfg)
	if err != nil {
		log.Fatalf("Can`t create store: %v", err)
	}
	defer newStore.MongoClient.Disconnect(myContext)
	collection := newStore.MongoClient.Database(cfg.MongoDb.DB).Collection(cfg.MongoDb.COLLECTION)

	classRepo := impl.NewClassRepository(collection)
	classService := impl2.NewClassService(classRepo)
	classController := controller.NewClassController(classService)

	server := api.NewServer([]api.Controller{classController})
	server.HandleRequests()
}
