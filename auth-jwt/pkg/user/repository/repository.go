package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	entities "gofiber-faafo/entities/user"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	Create(ctx context.Context, user *entities.User) error
	GetByEmail(c context.Context, email string) (entities.User, error)
	//Fetch() ([]entities.User, error)
	//GetByID(id string) (entities.User, error)
}

type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (ur *repository) Create(ctx context.Context, user *entities.User) error {
	collection := ur.Collection

	_, err := collection.InsertOne(ctx, user)

	return err
}

func (ur *repository) GetByEmail(c context.Context, email string) (entities.User, error) {
	collection := ur.Collection
	var user entities.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}
