package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/joho/godotenv"

	indexAPI "filtering/apis/index"
	processingAPI "filtering/apis/processing"
	"filtering/configuration"
	"filtering/utilities"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		BodyLimit:    2 * 1024 * 1024, // 2MB
		ErrorHandler: utilities.ErrorHandler,
	})

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = configuration.DEFAULT_ALLOWED_ORIGINS
	}

	app.Use(compress.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
	}))
	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
	}))
	app.Use(helmet.New())
	app.Use(limiter.New(limiter.Config{
		Expiration: 60 * time.Second,
		LimitReached: func(context *fiber.Ctx) error {
			return utilities.Response(utilities.ResponsePayloadStruct{
				Context: context,
				Info:    configuration.RESPONSE_MESSAGES.TooManyRequests,
				Status:  fiber.StatusTooManyRequests,
			})
		},
		Max: 30,
	}))
	app.Use(logger.New())

	indexAPI.Initialize(app)
	processingAPI.Initialize(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = configuration.DEFAULT_PORT
	}
	log.Fatal(app.Listen(":" + port))
}
