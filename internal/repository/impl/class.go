package impl

import (
	"context"
	"dttg/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

type ClassRepository struct {
	collection *mongo.Collection
}

func NewClassRepository(collection *mongo.Collection) *ClassRepository {
	return &ClassRepository{
		collection: collection,
	}
}

func (repo *ClassRepository) GetByName(ctx context.Context, name string) (*model.Class, error) {
	name = strings.Title(strings.ToLower(name))
	filter := bson.D{{"eng_name", name}}
	result := repo.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var class model.Class
	err := result.Decode(&class)
	if err != nil {
		return nil, err
	}

	return &class, nil
}
func (repo *ClassRepository) GetAll(ctx context.Context) ([]model.Class, error) {
	cursor, err := repo.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Читання результатів
	var classes []model.Class
	for cursor.Next(context.TODO()) {
		var result model.Class
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
