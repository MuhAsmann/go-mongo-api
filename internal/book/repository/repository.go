package repository

import (
	"context"
	"go-mongo-api/internal/book"
)

type BookRepository interface {
	FindAll(ctx context.Context) ([]book.Book, error)
	Create(ctx context.Context, book book.Book) (string, error)
}
