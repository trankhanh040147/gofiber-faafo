package usecase

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	databaserepository "gofiber-faafo/pkg/database/repository"
	"gofiber-faafo/pkg/index/model"
	"gofiber-faafo/pkg/index/repository"
	"gofiber-faafo/util"
)

type indexUsecase struct {
	repo   repository.IndexRepository
	dbRepo databaserepository.DatabaseRepository
}

type IndexUsecase interface {
	GetIndexes(ctx context.Context, dbName, collName string) ([]*model.Index, error)
	CreateIndex(ctx context.Context, dbName string, data model.Index) (*string, error)
}

func NewIndexUsecase(repo repository.IndexRepository, dbRepo databaserepository.DatabaseRepository) IndexUsecase {
	return &indexUsecase{
		repo:   repo,
		dbRepo: dbRepo,
	}
}

func (u *indexUsecase) GetIndexes(ctx context.Context, dbName, collName string) ([]*model.Index, error) {
	db, err := u.dbRepo.GetDBFromDBName(ctx, dbName)
	// check err not found
	if err != nil {
		return nil, err
	}

	client, err := util.GetNewDBConnection(db.Uri, db.DBName)
	if err != nil {
		return nil, errors.New("db connection error")
	}

	return u.repo.GetIndexes(ctx, client, collName)
}

func (u *indexUsecase) CreateIndex(ctx context.Context, dbName string, data model.Index) (*string, error) {
	db, err := u.dbRepo.GetDBFromDBName(ctx, dbName)
	if err != nil {
		return nil, err
	}

	client, err := util.GetNewDBConnection(db.Uri, db.DBName)
	if err != nil {
		return nil, errors.New("db connection error")
	}

	indexName, err := u.repo.CreateIndex(ctx, client, data)
	if err != nil {
		var convertErr mongo.CommandError
		if errors.As(err, &convertErr) {
			if convertErr.HasErrorCode(86) {
				return nil, errors.New("index conflict")
			}
		}
		return nil, err
	}

	return &indexName, nil
}
