package main

import (
	"blog_api/delivery/controller"
	"blog_api/delivery/router"
	"blog_api/infrastructure"
	"blog_api/repository"
	"blog_api/usecase"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	db := client.Database("blog_api")
	userCollection := db.Collection("users")
	postCollection := db.Collection("posts")

	postRepo := repository.NewPostRepository(postCollection)
	userRepo := repository.NewUserRepository(userCollection)
	
	ps := infrastructure.NewPasswordService()
	jwtService := infrastructure.NewJWTService(os.Getenv("JWT_SECRET"))
	postUsecase := usecase.NewPostUsecase(postRepo)
	userUsecase := usecase.NewUserUsecase(userRepo , ps , jwtService)
	postController := controller.NewPostController(postUsecase)
	userController := controller.NewUserController(userUsecase)
	middleware := infrastructure.NewMiddleware(jwtService)
	router := router.SetupRouter(postController , userController , *middleware)
	router.Run("localhost:3000")


}