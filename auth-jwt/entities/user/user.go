package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// todo: add not null for these fields: username, email, password, names

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"names" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	//Username  string             `json:"username" bson:"username"`
	//CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	//UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

const (
	CollectionUser = "users"
)
