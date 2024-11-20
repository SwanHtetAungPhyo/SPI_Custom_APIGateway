package services

import (
	"fmt"
	"github.com/SwanHtetAungPhyo/api/middleware"
	"math/rand"
	"net/http"
	"strings"
	"time"

	logger "github.com/SwanHtetAungPhyo/api/log"
	"github.com/SwanHtetAungPhyo/api/models"
	"github.com/SwanHtetAungPhyo/api/proxy"
	"github.com/SwanHtetAungPhyo/api/util"
	"github.com/gofiber/fiber/v2"
)

var logging = logger.GetLogger()

type GateWayServices struct{}

func NewGateWay() *GateWayServices {
	return &GateWayServices{}
}

func (receiver *GateWayServices) SetUpRoutes(app *fiber.App, config *models.GatewayConfig) {
	client := createHttpClient()
	rand.Seed(time.Now().UnixNano())

	for _, service := range config.Services {
		for _, route := range service.Routes {
			registerServiceRoutes(app, config, service, route, client, receiver)
		}
	}
}

func createHttpClient() *http.Client {
	return &http.Client{
		Timeout: 120 * time.Second,
		Transport: &http.Transport{
			ResponseHeaderTimeout: 120 * time.Second,
		},
	}
}

func registerServiceRoutes(app *fiber.App, config *models.GatewayConfig, service models.Service, route models.Route, client *http.Client, receiver *GateWayServices) {
	for _, path := range route.Path {
		logging.Println("Registering route:", fmt.Sprintf("/gate/%s%s", service.Name, path))
		app.All(fmt.Sprintf("/gate/%s%s", service.Name, path), func(ctx *fiber.Ctx) error {
			if shouldSkipJWT(ctx.Path()) {
				return makeProxyCall(ctx, service, client, receiver, config)
			} else {
				if middleware.JWTMiddleware(ctx) {
					return makeProxyCall(ctx, service, client, receiver, config)
				}
				return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Invalid or missing JWT token")
			}
		})
	}
}

func shouldSkipJWT(path string) bool {
	return strings.HasSuffix(path, "/login") || path == "/gate/services"
}

func makeProxyCall(ctx *fiber.Ctx, service models.Service, client *http.Client, receiver *GateWayServices, config *models.GatewayConfig) error {
	instanceAlgo := &util.InstanceAlgorithm{
		Algorithm: config.LoadBalancing,
	}
	currentInstance := util.GetCurrentInstance(instanceAlgo, service.Instance, ctx.IP())
	logging.Infof("Current instances is %v", currentInstance.Port)

	if currentInstance.Port == 0 {
		logging.Info("Fall into the fallback condition")
		currentInstance.Port = service.Instance[0]
	}

	lastPath := receiver.getLastPath(ctx.Path())
	pathSet := receiver.checkPathContain(service.Routes[0].Path)

	if pathSet[lastPath] {
		targetURL := fmt.Sprintf("%s:%d/%s%s", service.URL, currentInstance.Port, service.Leader, lastPath)
		logging.Println("Generated Target URL:", targetURL)
		return proxy.DoWithClient(ctx, targetURL, client)
	}

	ctx.Set("Information", lastPath)
	return ctx.Status(fiber.StatusBadRequest).Send([]byte(lastPath))
}

func (receiver *GateWayServices) getLastPath(path string) string {
	path = strings.TrimSuffix(path, "/")
	segments := strings.Split(path, "/")
	return "/" + segments[len(segments)-1]
}

func (receiver *GateWayServices) checkPathContain(paths []string) map[string]bool {
	set := make(map[string]bool)
	for _, item := range paths {
		set[item] = true
	}
	return set
}
