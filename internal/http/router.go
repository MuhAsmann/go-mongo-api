package http

import (
	bookHandler "go-mongo-api/internal/book/handler"
	userHandler "go-mongo-api/internal/user/handler"
	"net/http"
)

type Router struct {
	BookHandler *bookHandler.BookHandler
	UserHandler *userHandler.UserHandler
}

func NewRouter(bh *bookHandler.BookHandler, uh *userHandler.UserHandler) *Router {
	return &Router{
		BookHandler: bh,
		UserHandler: uh,
	}
}

func (r *Router) InitRoutes() {
	http.HandleFunc("/books", r.BookHandler.GetBooks)
	http.HandleFunc("/books/create", r.BookHandler.CreateBook)

	http.HandleFunc("/users", r.UserHandler.GetUserByID)
}
