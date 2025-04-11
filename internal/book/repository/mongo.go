package repository

import (
	"context"
	"go-mongo-api/internal/book"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo(col *mongo.Collection) BookRepository {
	return &mongoRepo{collection: col}
}

func (r *mongoRepo) FindAll(ctx context.Context) ([]book.Book, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []book.Book
	for cursor.Next(ctx) {
		var b book.Book
		cursor.Decode(&b)
		books = append(books, b)
	}
	return books, nil
}

func (r *mongoRepo) Create(ctx context.Context, b book.Book) (string, error) {
	res, err := r.collection.InsertOne(ctx, b)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
