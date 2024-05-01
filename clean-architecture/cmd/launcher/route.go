package launcher

import (
	"gofiber-faafo/api/handlers"
	"gofiber-faafo/pkg/book"
	databasehandler "gofiber-faafo/pkg/database/handler"
	databaserepository "gofiber-faafo/pkg/database/repository"
	databaseusecase "gofiber-faafo/pkg/database/usecase"
	indexhandler "gofiber-faafo/pkg/index/handler"
	indexrepository "gofiber-faafo/pkg/index/repository"
	indexusecase "gofiber-faafo/pkg/index/usecase"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoute(app *fiber.App, db *mongo.Database) {
	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
	bookService := book.NewService(bookRepo)

	databaseRepo := databaserepository.NewDatabaseRepository(db, "databases")
	databaseUC := databaseusecase.NewDatabaseUsecase(databaseRepo)
	databaseHandler := databasehandler.NewDatabaseHandler(databaseUC)

	indexRepo := indexrepository.NewIndexRepository(db, "books")
	indexUC := indexusecase.NewIndexUsecase(indexRepo, databaseRepo)
	indexHandler := indexhandler.NewIndexHandler(indexUC)

	api := app.Group("/api")

	books := api.Group("/books")
	{
		books.Get("/", handlers.GetBooks(bookService))
		books.Post("/", handlers.AddBook(bookService))
		books.Post("/many", handlers.AddBooks(bookService))
		books.Put("/", handlers.UpdateBook(bookService))
		books.Delete("/", handlers.RemoveBook(bookService))
	}

	databases := api.Group("/databases")
	{
		databases.Get("/", databaseHandler.GetDatabases())
		databases.Post("/", databaseHandler.CreateDatabase())
		indexes := databases.Group("/:db_name/indexes")
		{
			indexes.Get("/", indexHandler.GetIndexes())
			indexes.Post("/", indexHandler.CreateIndex())
		}
	}

}
