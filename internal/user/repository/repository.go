package repository

import (
	"context"
	model "go-mongo-api/internal/user"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (model.User, error)
}
