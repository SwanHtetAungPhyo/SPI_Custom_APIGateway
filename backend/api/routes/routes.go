package routes

import (
	"github.com/SwanHtetAungPhyo/api/handler"
	"github.com/SwanHtetAungPhyo/api/models"
	"github.com/SwanHtetAungPhyo/api/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesForAPP(app *fiber.App, gatewayServices *services.GateWayServices) {
	gateway := app.Group("/gate/")
	gatewayHandler := handler.GateWayHandler{}

	gateway.Get("/services", gatewayHandler.Services)
	gatewayServices.SetUpRoutes(app, models.Configuration())
}
