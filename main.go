package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gofiber-faafo/api/routes"
	"gofiber-faafo/pkg/book"
	"log"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")

	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
	bookService := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})

	api := app.Group("/api")
	routes.BookRouter(api, bookService)

	defer cancel()
	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://127.0.0.1:27017/fiber?directConnection=true").SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("books")
	return db, cancel, nil
}