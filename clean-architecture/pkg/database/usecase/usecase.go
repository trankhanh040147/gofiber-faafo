package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gofiber-faafo/pkg/database/model"
	"gofiber-faafo/pkg/database/repository"
)

type databaseUsecase struct {
	repo repository.DatabaseRepository
}

type DatabaseUsecase interface {
	GetDatabases(ctx context.Context) ([]*model.Database, error)
	CreateDatabase(ctx context.Context, data model.Database) (primitive.ObjectID, error)
}

func NewDatabaseUsecase(repo repository.DatabaseRepository) DatabaseUsecase {
	return &databaseUsecase{
		repo: repo,
	}
}

func (u *databaseUsecase) GetDatabases(ctx context.Context) ([]*model.Database, error) {
	return u.repo.GetDatabases(ctx)
}

func (u *databaseUsecase) CreateDatabase(ctx context.Context, data model.Database) (primitive.ObjectID, error) {
	return u.repo.CreateDatabase(ctx, data)
}
