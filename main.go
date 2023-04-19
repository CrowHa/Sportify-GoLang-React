package main

import (
	"log"

	"Go-Sportify/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "*",
	}))

	app.Post("/login", services.Login)

	if err := app.Listen(":3030"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
