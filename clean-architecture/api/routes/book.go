package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiber-faafo/api/handlers"
	"gofiber-faafo/pkg/book"
)

// BookRouter is the Router for GoFiber App
func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/books", handlers.GetBooks(service))
	app.Post("/books", handlers.AddBook(service))
	app.Post("/books-many", handlers.AddBooks(service))
	app.Put("/books", handlers.UpdateBook(service))
	app.Delete("/books", handlers.RemoveBook(service))
}
