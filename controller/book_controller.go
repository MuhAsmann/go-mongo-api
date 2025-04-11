package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go-mongo-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookController struct {
	Collection *mongo.Collection
}

func (bc *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := bc.Collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var books []model.Book
	for cursor.Next(ctx) {
		var book model.Book
		if err := cursor.Decode(&book); err == nil {
			books = append(books, book)
		}
	}

	json.NewEncoder(w).Encode(books)
}

func (bc *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := bc.Collection.InsertOne(ctx, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book.ID = res.InsertedID.(primitive.ObjectID).Hex()
	json.NewEncoder(w).Encode(book)
}
