package main

import (
	"github.com/SwanHtetAungPhyo/api/handler"
	"github.com/SwanHtetAungPhyo/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,DELETE",
		AllowHeaders: "application/json",
	}))

	app.Use(middleware.MiddleMan{}.ResponseTimeRuler)

	gateway := app.Group("/gate/")
	gatewayHandler := handler.GateWayHandler{}

	gateway.Get("/services", gatewayHandler.Services)
	log.Fatal(app.Listen(":8081"))
}
