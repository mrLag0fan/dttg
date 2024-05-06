package impl

import (
	"context"
	"dttg/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

type RaceRepository struct {
	collection *mongo.Collection
}

func NewRaceRepository(collection *mongo.Collection) *RaceRepository {
	return &RaceRepository{
		collection: collection,
	}
}

func (repo *RaceRepository) GetByName(ctx context.Context, name string) (*interface{}, error) {
	name = strings.Title(strings.ToLower(name))
	filter := bson.D{{"eng_name", name}}
	result := repo.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var class model.Race
	err := result.Decode(&class)
	if err != nil {
		return nil, err
	}

	var obj interface{} = class
	return &obj, nil
}
func (repo *RaceRepository) GetAll(ctx context.Context) ([]interface{}, error) {
	cursor, err := repo.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Читання результатів
	var classes []interface{}
	for cursor.Next(context.TODO()) {
		var result model.Race
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		classes = append(classes, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return classes, nil
}
