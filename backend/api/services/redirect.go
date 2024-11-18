package services

import (
	"fmt"
	logger "github.com/SwanHtetAungPhyo/api/log"
	"github.com/SwanHtetAungPhyo/api/middleware"
	"github.com/SwanHtetAungPhyo/api/models"
	"github.com/SwanHtetAungPhyo/api/proxy"
	"github.com/SwanHtetAungPhyo/api/util"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var logging = logger.GetLogger()

type GateWayServices struct{}

func NewGateWay() *GateWayServices {
	return &GateWayServices{}
}
func (receiver *GateWayServices) SetUpRoutes(app *fiber.App, config *models.GatewayConfig) {
	for _, service := range config.Services {
		for _, route := range service.Routes {
			requestTimeout, _ := time.ParseDuration(route.Timeout)
			writeTimeout := requestTimeout
			for _, path := range route.Path {
				logging.Println("Registering route:", fmt.Sprintf("/gate/%s%s", service.Name, path))
				app.All(fmt.Sprintf("/gate/%s%s", service.Name, path), func(ctx *fiber.Ctx) error {
					if ctx.Path() != "/login" && ctx.Path() != "/gate/services" && !strings.HasPrefix(ctx.Path(), "/gate/") {
						return middleware.JWTMiddleware(ctx)
					}

					client := &http.Client{
						Timeout: requestTimeout,
						Transport: &http.Transport{
							ResponseHeaderTimeout: writeTimeout,
						},
					}

					instanceAlgo := &util.InstanceAlgorithm{
						Algorithm: config.LoadBalancing,
					}
					currentInstance := util.GetCurrentInstance(instanceAlgo, service.Instance)

					logging.Printf("Port: %d, Instance: %+v", currentInstance.Port, currentInstance)
					if currentInstance.Port == 0 {
						currentInstance.Port = service.Instance[0]
					}

					targetURL := fmt.Sprintf("%s:%d/%s", service.URL, currentInstance.Port, service.Leader)
					logging.Println("Generated Target URL:", targetURL)

					return proxy.DoWithClient(ctx, targetURL, client)
				})

			}
		}
	}
}

func (receiver *GateWayServices) getRandomInstance(instances []int) int {
	rand.Seed(time.Now().UnixNano())
	return instances[rand.Intn(len(instances))]
}
