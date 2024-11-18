package main

import "github.com/gofiber/fiber/v2"

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	app := fiber.New()

	app.Get("/user/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello, World!"))
	})

	app.Post("/user/", func(c *fiber.Ctx) error {
		var user User
		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(400).Send([]byte("Bad Request"))
		}
		return c.Status(201).JSON(user)
	})
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
