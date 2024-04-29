package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"gofiber-faafo/pkg/user/repository"
	"gofiber-faafo/pkg/user/usecase"
	"time"
)

func NewSignupRouter(timeout time.Duration, db mongo.Database, group *fiber.Router) {
	ur := repository.NewRepo()
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
	}
	group.POST("/signup", sc.Signup)
}

// UserRoute is the Router for GoFiber App
//func UserHandler(app fiber.Router, UC usecase
//) {
//	app.Get("/books", handlers.GetBooks(service))
//	app.Post("/books", handlers.AddBook(service))
//	app.Put("/books", handlers.UpdateBook(service))
//	app.Delete("/books", handlers.RemoveBook(service))
//}
