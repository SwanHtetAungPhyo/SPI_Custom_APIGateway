package main

import (
	"github.com/SwanHtetAungPhyo/api/middleware"
	"github.com/SwanHtetAungPhyo/api/routes"
	"github.com/SwanHtetAungPhyo/api/services"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

type myCustomStorage struct {
}

// @Author: Swan Htet Aung Phyo
// @StartDate: Nov 12 2024
// @MainTechnology: go fiber
func main() {
	log.SetOutput(os.Stdout)
	app := fiber.New(
		fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		})
	middleman := middleware.NewMiddleMan()
	middleman.SetupMiddlewares(app)
	gatewayServices := services.NewGateWay()
	routes.SetupRoutesForAPP(app, gatewayServices)
	log.Fatal(app.Listen(":8081"))
}
