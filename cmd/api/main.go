package main

import (
	config "go-mongo-api/internal"
	bookHandler "go-mongo-api/internal/book/handler"
	bookRepo "go-mongo-api/internal/book/repository"
	bookUsecase "go-mongo-api/internal/book/usecase"
	appHttp "go-mongo-api/internal/http"
	userHandler "go-mongo-api/internal/user/handler"
	userRepo "go-mongo-api/internal/user/repository"
	userUsecase "go-mongo-api/internal/user/usecase"
	"log"
	"net/http"
)

func main() {
	client, err := config.InitMongo("mongodb://root:example@mongodb:27017")
	if err != nil {
		log.Fatal(err)
	}

	col := client.Database("bookstore").Collection("books")
	bookRepo := bookRepo.NewMongoRepo(col)
	bookUsecase := bookUsecase.New(bookRepo)
	bookHandler := bookHandler.New(bookUsecase)

	userCol := client.Database("bookstore").Collection("users")
	userRepository := userRepo.NewMongoRepo(userCol)
	userUC := userUsecase.New(userRepository)
	userHandler := userHandler.NewUserHandler(userUC)

	// Routing
	router := appHttp.NewRouter(bookHandler, userHandler)
	router.InitRoutes()

	log.Println("✅ Server started at :8000")
	http.ListenAndServe(":8000", nil)
}
