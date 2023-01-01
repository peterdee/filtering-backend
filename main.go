package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"filtering/configuration"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load the .env file!")
	}

	// TODO:
	app := fiber.New()

	// TODO: middlewares (helmet, limiter)
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
	}))
	app.Use(logger.New())

	// TODO: API for image processing

	port := os.Getenv("PORT")
	if port == "" {
		port = configuration.DEFAULT_PORT
	}
	log.Fatal(app.Listen(":" + port))
}
