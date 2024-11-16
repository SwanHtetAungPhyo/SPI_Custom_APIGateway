package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"log"
	"os"
	"strconv"
	"time"
)

type MiddleMan struct {
}

func NewMiddleMan() *MiddleMan {
	return &MiddleMan{}
}

func (m *MiddleMan) SetupMiddlewares(app *fiber.App) {

	app.Use(healthcheck.New())

	app.Use(m.ResponseTimeRuler)

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendFile("./toofast.html")
		},
		Storage: nil,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,DELETE",
		AllowHeaders: "application/json,Authorization",
	}))

	app.Use(earlydata.New(earlydata.Config{
		Error: fiber.ErrTooEarly,
	}))

	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output:     os.Stdout,
		TimeFormat: time.RFC3339Nano,
		TimeZone:   "Europe/Warsaw",
		CustomTags: map[string]logger.LogFunc{
			"request_id": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(c.Locals("requestid").(string))
			},
		},
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "SPI GateWay"}))
}
func (m *MiddleMan) ResponseTimeRuler(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)
	durationInMil := duration.Microseconds()

	durationInSec := duration.Seconds()
	key := fmt.Sprintf("Response time from %s is %v", c.OriginalURL(), durationInSec)
	log.Println(key)
	c.Set(key, strconv.FormatInt(durationInMil, 10))
	return err
}
