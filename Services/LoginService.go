package services

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"Go-Sportify/dbconnection"
	"Go-Sportify/model"
)

func Login(c *fiber.Ctx) error {
	form := new(LoginForm)
	if err := c.BodyParser(form); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	db, err := dbconnection.Connection()
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// var users []string
	// result := db.Model(&model.User{}).Pluck("username", &users)
	// if result.Error != nil {
	// 	log.Println(result.Error)
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

	user := model.User{}
	if err := db.Where(&model.User{Username: form.Username, Password: form.Password}).First(&user).Error; err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}

// form := new(LoginForm)
// if err := c.BodyParser(form); err != nil {
// 	return err
// }

// // Validate login credentials (for example, check against database)
// if form.Username != "admin" || form.Password != "password" {
// 	c.SendString("Sai Pass Or TK")
// 	return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
// }

// // Return success response
// return c.JSON(fiber.Map{
// 	"message": "Login successful",
// })

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
