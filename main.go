package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
	"github.com/thitiph0n/go-url-shortener/handlers"
	"github.com/thitiph0n/go-url-shortener/repositories"
	"github.com/thitiph0n/go-url-shortener/services"
)

func main() {

	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Print(".env file not found")
	}

	ctx := context.Background()

	firestoreClient := initFirestore(ctx)
	defer firestoreClient.Close()

	linkRepo := repositories.NewLinkRepositoryFirestore(ctx, firestoreClient)
	linkService := services.NewLinkService(linkRepo)
	linkHandler := handlers.NewLinkHandler(linkService)

	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.HandleError,
	})

	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("APP_ALLOW_ORIGIN"),
	}))

	app.Get("/links", linkHandler.GetLinks)
	app.Get("/links/:linkId", linkHandler.GetLinkById)
	app.Post("/links", linkHandler.CreateLink)
	app.Get("/reslove/:linkId", linkHandler.ResloveLink)

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}

}

func initFirestore(ctx context.Context) *firestore.Client {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
