package usecase

import (
	"context"
	model "go-mongo-api/internal/user"
	"go-mongo-api/internal/user/repository"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (model.User, error)
}

type usecase struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) UserUsecase {
	return &usecase{repo}
}

func (uc *usecase) GetByID(ctx context.Context, id string) (model.User, error) {
	return uc.repo.FindByID(ctx, id)
}
