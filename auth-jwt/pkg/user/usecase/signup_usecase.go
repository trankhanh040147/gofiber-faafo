package usecase

import (
	"context"
	"gofiber-faafo/domain"
	entities "gofiber-faafo/entities/user"
	"gofiber-faafo/pkg/user/repository"
	"time"
)

type signupUsecase struct {
	userRepository repository.Repository
	contextTimeout time.Duration
}

func (su *signupUsecase) Create(c context.Context, user *entities.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (entities.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func NewSignupUsecase(userRepository repository.Repository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
