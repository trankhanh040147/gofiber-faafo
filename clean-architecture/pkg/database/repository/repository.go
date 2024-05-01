package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gofiber-faafo/pkg/database/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type databaseRepository struct {
	db         *mongo.Database
	collection string
}

type DatabaseRepository interface {
	GetDatabases(ctx context.Context) ([]*model.Database, error)
	GetDBFromDBName(ctx context.Context, dbName string) (*model.Database, error)
	CreateDatabase(ctx context.Context, data model.Database) (primitive.ObjectID, error)
}

func NewDatabaseRepository(db *mongo.Database, collection string) DatabaseRepository {
	return &databaseRepository{
		db:         db,
		collection: collection,
	}
}

func (r *databaseRepository) GetDatabases(ctx context.Context) ([]*model.Database, error) {
	var databases []*model.Database
	coll := r.db.Collection(r.collection)
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var database model.Database
		_ = cursor.Decode(&database)
		databases = append(databases, &database)
	}
	return databases, nil
}

func (r *databaseRepository) CreateDatabase(ctx context.Context, data model.Database) (primitive.ObjectID, error) {
	coll := r.db.Collection(r.collection)
	insertResult, err := coll.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return insertResult.InsertedID.(primitive.ObjectID), nil
}

func (r *databaseRepository) GetDBFromDBName(ctx context.Context, dbName string) (*model.Database, error) {
	coll := r.db.Collection(r.collection)
	var database model.Database
	err := coll.FindOne(ctx, bson.D{{"name", dbName}}).Decode(&database)
	if err != nil {
		return nil, err
	}
	return &database, nil
}
