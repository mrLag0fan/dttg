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
	classCollection := newStore.MongoClient.Database(cfg.MongoDb.DB).Collection(cfg.MongoDb.CLASS_COLLECTION)
	spellCollection := newStore.MongoClient.Database(cfg.MongoDb.DB).Collection(cfg.MongoDb.SPELL_COLLECTION)

	classRepo := impl.NewClassRepository(classCollection)
	classService := impl2.NewClassService(classRepo)
	classController := controller.NewClassController(classService)

	spellRepo := impl.NewSpellRepository(spellCollection)
	spellService := impl2.NewSpellService(spellRepo)
	spellController := controller.NewSpellController(spellService)

	server := api.NewServer([]api.Controller{classController, spellController})
	server.HandleRequests()
}
