package services

import (
	"fmt"
	"github.com/SwanHtetAungPhyo/api/models"
	"github.com/SwanHtetAungPhyo/api/proxy"
	"github.com/SwanHtetAungPhyo/api/util"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

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
				app.All(fmt.Sprintf("/gate/%s%s", service.Name, path), func(ctx *fiber.Ctx) error {
					id := ctx.Params("id")

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
					log.Println("Selected Instance:", currentInstance)

					targetURL := fmt.Sprintf("%s:%d/%s", service.URL, currentInstance.Port, service.Leader)

					if id != "" {
						if ind, err := strconv.Atoi(id); err == nil {
							targetURL = fmt.Sprintf("%s/%v", targetURL, ind)
						} else {
							log.Println("Invalid ID format:", err)
						}
					}

					log.Println("Target URL:", targetURL)
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
