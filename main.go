package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/thitiph0n/go-url-shortener/handlers"
	"github.com/thitiph0n/go-url-shortener/repositories"
	"github.com/thitiph0n/go-url-shortener/services"
	"google.golang.org/api/option"
)

func main() {

	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal(err)
	}

	firestoreClient := initFirestore()
	defer firestoreClient.Close()

	linkRepo := repositories.NewLinkRepositoryFirestore(firestoreClient)
	linkService := services.NewLinkService(linkRepo)
	linkHandler := handlers.NewLinkHandler(linkService)

	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.HandleError,
	})

	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Get("/links", linkHandler.GetLinks)
	app.Get("/links/:linkId", linkHandler.GetLinkById)
	app.Post("/links", linkHandler.CreateLink)
	app.Get("/reslove/:linkId", linkHandler.ResloveLink)

	if err := app.Listen(os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err)
	}

}

func initFirestore() *firestore.Client {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./config/thitiph0n-go-url-shortener-e534e5e27489.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
