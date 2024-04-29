package domain

import (
	"context"
	entities "gofiber-faafo/entities/user"
)

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	//AccessToken  string `json:"accessToken"`
	//RefreshToken string `json:"refreshToken"`
	Status bool `json:"status"`
}

type SignupUsecase interface {
	Create(c context.Context, user *entities.User) error
	GetUserByEmail(c context.Context, email string) (entities.User, error)
	//CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error)
	//CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error)
}
