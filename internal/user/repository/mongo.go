package repository

import (
	"context"
	model "go-mongo-api/internal/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	col *mongo.Collection
}

func NewMongoRepo(col *mongo.Collection) UserRepository {
	return &mongoRepo{col}
}

func (r *mongoRepo) FindByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	err := r.col.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	return user, err
}
