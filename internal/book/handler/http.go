package handler

import (
	"context"
	"encoding/json"
	"go-mongo-api/internal/book"
	"go-mongo-api/internal/book/usecase"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookHandler struct {
	Usecase usecase.BookUsecase
}

func New(uc usecase.BookUsecase) *BookHandler {
	return &BookHandler{Usecase: uc}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := h.Usecase.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var b book.Book
	json.NewDecoder(r.Body).Decode(&b)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := h.Usecase.Add(ctx, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "invalid ID format", http.StatusInternalServerError)
		return
	}
	b.ID = objectID
	json.NewEncoder(w).Encode(b)
}
