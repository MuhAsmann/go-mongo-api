package handler

import (
	"go-mongo-api/internal/user/usecase"
	"net/http"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting user by ID
}
