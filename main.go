package main

import (
	"go-mongo-api/config"
	"go-mongo-api/controller"
	"log"
	"net/http"
)

func main() {
	config.Connect()

	bookCollection := config.GetCollection("books")
	bookCtrl := &controller.BookController{Collection: bookCollection}

	http.HandleFunc("/books", bookCtrl.GetBooks)
	http.HandleFunc("/books/create", bookCtrl.CreateBook)

	log.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
