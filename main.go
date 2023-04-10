package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "*",
	}))

	// Handle POST request to login endpoint
	app.Post("/login", func(c *fiber.Ctx) error {
		// Parse form data
		form := new(LoginForm)
		if err := c.BodyParser(form); err != nil {
			return err
		}

		// Validate login credentials (for example, check against database)
		if form.Username != "admin" || form.Password != "password" {
			c.SendString("Sai Pass Or TK")
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
		}

		// Return success response
		return c.JSON(fiber.Map{
			"message": "Login successful",
		})
	})
	app.Get("/He", func(c *fiber.Ctx) error {
		return c.SendString("HeLLO")
	})

	// Start server
	app.Listen(":3030")
}
