package usecase

import (
	"context"
	"go-mongo-api/internal/book"
	"go-mongo-api/internal/book/repository"
)

type BookUsecase interface {
	GetAll(ctx context.Context) ([]book.Book, error)
	Add(ctx context.Context, b book.Book) (string, error)
}

type usecase struct {
	repo repository.BookRepository
}

func New(repo repository.BookRepository) BookUsecase {
	return &usecase{repo: repo}
}

func (u *usecase) GetAll(ctx context.Context) ([]book.Book, error) {
	return u.repo.FindAll(ctx)
}

func (u *usecase) Add(ctx context.Context, b book.Book) (string, error) {
	return u.repo.Create(ctx, b)
}
