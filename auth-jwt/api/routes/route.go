package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"gofiber-faafo/api/controller"
	"gofiber-faafo/pkg/user/repository"
	"gofiber-faafo/pkg/user/usecase"
)

func SetupRoute(app fiber.Router, db *mongo.Database) {
	userCollection := db.Collection("users")
	userRepo := repository.NewRepo(userCollection)
	//userUsecase := usecase.NewUserUseCase(userRepo)
	//userHandler = handler.NewUserHandler(userUsecase)
	//signupUC := usecase.NewSignupUsecase
	//profileUC := usecase.ProfileUsecase

	//signupController := controller.SignupController{SignupUsecase: signupUC}
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(userRepo, 0),
		//Env:           env,
	}

	//route: User
	//app.Post("/signup", userHandler.Signup(userUsecase))
	app.Get("/signup", sc.Signup)
}
