package services

import (
	"fmt"
	"github.com/SwanHtetAungPhyo/api/middleware"
	"github.com/SwanHtetAungPhyo/api/models"
	"github.com/SwanHtetAungPhyo/api/proxy"
	"github.com/SwanHtetAungPhyo/api/util"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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
				log.Println("Registering route:", fmt.Sprintf("/gate/%s%s", service.Name, path))
				receiver.registerRoute(app, service, config, path, requestTimeout, writeTimeout)
			}
		}
	}
}

func (receiver *GateWayServices) registerRoute(app *fiber.App, service models.Service, config *models.GatewayConfig, path string, requestTimeout, writeTimeout time.Duration) {
	app.All(fmt.Sprintf("/gate/jwt/%s%s", service.Name, path), func(ctx *fiber.Ctx) error {
		if ctx.Path() != "/gate/services" && !strings.HasPrefix(ctx.Path(), "/gate/") {
			log.Println("JWT Middleware applied to", ctx.Path())
			return middleware.JWTMiddleware(ctx)
		}

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
		if currentInstance.Port == 0 {
			log.Println("Error: Selected instance has port 0. Defaulting to a fallback port.")
			currentInstance.Port = service.Instance[0]
		}

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

func (receiver *GateWayServices) getRandomInstance(instances []int) int {
	rand.Seed(time.Now().UnixNano())
	return instances[rand.Intn(len(instances))]
}
