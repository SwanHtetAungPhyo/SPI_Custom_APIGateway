package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users  = make(map[int]User)
	mu     sync.Mutex
	nextID = 1
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8081",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "application/json",
	}))

	app.Post("/user", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}
		mu.Lock()
		user.ID = nextID
		nextID++
		users[user.ID] = user
		mu.Unlock()
		return c.Status(fiber.StatusCreated).JSON(user)
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		mu.Lock()
		defer mu.Unlock()
		var allUsers []User
		for _, user := range users {
			allUsers = append(allUsers, user)
		}
		return c.JSON(allUsers)
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user User
		mu.Lock()
		defer mu.Unlock()
		for _, u := range users {
			if fmt.Sprint(u.ID) == id {
				user = u
				break
			}
		}
		if user.ID == 0 {
			return c.Status(fiber.StatusNotFound).SendString("User not found")
		}
		return c.JSON(user)
	})

	app.Put("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}
		mu.Lock()
		defer mu.Unlock()
		for i, u := range users {
			if fmt.Sprint(u.ID) == id {
				users[i] = user
				user.ID = u.ID
				return c.JSON(user)
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	})

	app.Delete("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		mu.Lock()
		defer mu.Unlock()
		for i, u := range users {
			if fmt.Sprint(u.ID) == id {
				delete(users, i)
				return c.SendString("User deleted successfully")
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	})

	go func() {
		if err := app.Listen(":3001"); err != nil {
			log.Fatalf("Error starting server on port 4001: %v", err)
		}
	}()
	go func() {
		if err := app.Listen(":3002"); err != nil {
			log.Fatalf("Error starting server on port 4002: %v", err)
		}
	}()
	go func() {
		if err := app.Listen(":3003"); err != nil {
			log.Fatalf("Error starting server on port 4003: %v", err)
		}
	}()

	select {}
}
