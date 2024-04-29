package launcher

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDBConnection() (*mongo.Database, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://127.0.0.1:27017/fiber?directConnection=true").SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		log.Fatal("Database Connection Error $s", err)
		fmt.Println("Database connection success!")
		return nil, nil
	}
	db := client.Database("books")
	return db, cancel
}
