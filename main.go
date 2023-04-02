package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string
	Role     string
}

func main() {

	app := fiber.New()

	app.Get("/post", HandleGetPost)
	app.Get("/post/manage", OnlyAdmin(HandleGetPostManager))

	log.Fatal(app.Listen(":1998"))
}

func OnlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := GetUserFromDB()
		if user.Role != "admin" {
			return c.SendStatus(http.StatusUnauthorized)
		}

		return fn(c)
	}
}

func GetUserFromDB() User {
	return User{
		Username: "Daniel",
		Role:     "user",
	}
}

func HandleGetPost(c *fiber.Ctx) error {
	return c.JSON("Handle something Here")
}

func HandleGetPostManager(c *fiber.Ctx) error {
	return c.JSON("The Admin Page")
}
