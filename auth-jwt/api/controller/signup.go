package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gofiber-faafo/domain"
	entities "gofiber-faafo/entities/user"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	//Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *fiber.Ctx) {
	var request domain.SignupRequest

	err := c.BodyParser(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.Status(http.StatusConflict)
		c.JSON(domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.JSON(domain.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := entities.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.JSON(domain.ErrorResponse{Message: err.Error()})
		return
	}

	//accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	//	return
	//}
	//
	//refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	//	return
	//}

	//signupResponse := domain.SignupResponse{
	//	AccessToken:  accessToken,
	//	RefreshToken: refreshToken,
	//}

	signupResponse := domain.SignupResponse{
		Status: true,
	}

	c.Status(http.StatusOK)
	c.JSON(signupResponse)
}
