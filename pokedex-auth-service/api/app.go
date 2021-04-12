package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gus-messagi/pokedex-api/pokedex-auth-service/api/routes"
	"github.com/gus-messagi/pokedex-api/pokedex-auth-service/pkg/user"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	db, err := DatabaseConnection()

	if err != nil {
		log.Fatal("Database connection error $s", err)
	}

	userCollection := db.Collection("users")
	userRepo := user.NewRepo(userCollection)
	userService := user.NewService(userRepo)

	app := fiber.New()

	api := app.Group("/api")

	routes.AuthRouter(api, userService)

	_ = app.Listen(":8080")
}

func DatabaseConnection() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(envVariable("MONGO_URL")))

	if err != nil {
		return nil, err
	}

	db := client.Database("users")

	return db, nil
}

func envVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
