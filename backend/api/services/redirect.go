package services

import (
	"fmt"
	"github.com/SwanHtetAungPhyo/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"math/rand"
	"time"
)

type GateWayServices struct {
}

func (receiver GateWayServices) SetUpRoutes(app *fiber.App, config *models.GatewayConfig) {
	for _, service := range config.Services {
		for _, route := range service.Routes {
			app.All("/gate/"+service.Name+route.Path, func(ctx *fiber.Ctx) error {
				currentInstance := receiver.getRandomInstance(service.Instance)
				targetURL := fmt.Sprintf("%s:%d%s", service.URL, currentInstance, route.Path)
				return proxy.Do(ctx, targetURL)
			})
		}
	}
}

func (receiver *GateWayServices) getRandomInstance(instances []int) int {
	rand.Seed(time.Now().UnixNano())
	return instances[rand.Intn(len(instances))]
}
