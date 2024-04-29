package launcher

import (
	"gofiber-faafo/api/routes"
	"gofiber-faafo/pkg/book"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoute(app *fiber.App, db *mongo.Database) {
	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
	bookService := book.NewService(bookRepo)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})

	api := app.Group("/api")
	routes.BookRouter(api, bookService)
}
