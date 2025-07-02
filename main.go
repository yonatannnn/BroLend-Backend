package main

import (
	"brolend/controller"
	"brolend/infrastructure"
	"brolend/repository"
	router "brolend/router"
	"brolend/usecase"
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	dbName := os.Getenv("MONGO_DB")
	userCollection := client.Database(dbName).Collection("users")
	userRepo := repository.NewUserRepository(userCollection, context.TODO())
	passwordService := infrastructure.NewPasswordService()
	jwtService := infrastructure.NewJwtService(os.Getenv("JWT_SECRET"))
	fmt.Println("JWT_SECRET", os.Getenv("JWT_SECRET"))
	userUsecase := usecase.NewUserUsecase(userRepo, passwordService, jwtService)
	userController := controller.NewUserController(userUsecase)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r := router.SetupRouter(userController)
	r.Run(":" + port)
}
