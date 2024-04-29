package usecase

import "gofiber-faafo/pkg/user/repository"

type accountUseCase struct {
	userRepo repository.Repository
}

func NewAccountUseCase(userRepo repository.Repository) *accountUseCase {
	return &accountUseCase{
		userRepo: userRepo,
	}
}
