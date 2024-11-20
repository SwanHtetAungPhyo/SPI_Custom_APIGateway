package main

import (
	logger "github.com/SwanHtetAungPhyo/api/log"
	"github.com/SwanHtetAungPhyo/api/middleware"
	"github.com/SwanHtetAungPhyo/api/models"
	"github.com/SwanHtetAungPhyo/api/routes"
	"github.com/SwanHtetAungPhyo/api/services"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"log"
	_ "net/http/pprof"
)

// @Author: Swan Htet Aung Phyo
// @StartDate: Nov 12 2024
// @MainTechnology: go fiber
func main() {
	logger.SetupLogger()
	log.Println("Starting the application...")
	app := fiber.New(
		fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		})
	gatewayServices := services.NewGateWay()
	config := models.Configuration()
	middleman := middleware.NewMiddleMan(config.BlackSpace)
	middleman.SetupMiddlewares(app)


	routes.SetupRoutesForAPP(app, gatewayServices)
	log.Fatal(app.Listen(":8081"))

}
