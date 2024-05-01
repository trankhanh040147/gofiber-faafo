package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofiber-faafo/pkg/index/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type indexRepository struct {
	db         *mongo.Database
	collection string
}

type IndexRepository interface {
	GetIndexes(ctx context.Context, client *mongo.Database, collName string) ([]*model.Index, error)
	CreateIndex(ctx context.Context, client *mongo.Database, data model.Index) (string, error)
}

func NewIndexRepository(db *mongo.Database, collection string) IndexRepository {
	return &indexRepository{
		db:         db,
		collection: collection,
	}
}

func (r *indexRepository) GetIndexes(ctx context.Context, client *mongo.Database, collName string) ([]*model.Index, error) {
	coll := client.Collection(collName)
	cursor, err := coll.Indexes().List(ctx)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var indexes []interface{}
	for cursor.Next(ctx) {
		var index interface{}
		if err := cursor.Decode(&index); err != nil {
			return nil, err
		}
		indexes = append(indexes, index)
	}

	indexModels := make([]*model.Index, len(indexes))
	for key, value := range indexes {
		indexModels[key] = model.ConvertFromMongoIndex(value.(primitive.D))
	}

	return indexModels, nil
}

func (r *indexRepository) CreateIndex(ctx context.Context, client *mongo.Database, data model.Index) (string, error) {
	coll := client.Collection(data.Collection)

	indexKeys := bson.D{}
	for key, value := range data.Keys {
		k := bson.E{Key: key, Value: value}
		indexKeys = append(indexKeys, k)
	}

	indexModel := mongo.IndexModel{
		Keys:    indexKeys,
		Options: options.Index().SetUnique(data.IsUnique),
	}

	fmt.Println(indexModel.Keys)

	result, err := coll.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return "", err
	}

	return result, nil
}
