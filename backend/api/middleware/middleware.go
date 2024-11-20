package middleware

import (
	"fmt"
	logger "github.com/SwanHtetAungPhyo/api/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/oschwald/maxminddb-golang"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var logging = logger.GetLogger()
var blockedAgents = map[string]bool{
	"badbot":                               true,
	"scraper":                              true,
	"curl":                                 true,
	"wget":                                 true,
	"python-urllib":                        true,
	"postmanruntime":                       true,
	"ahrefsbot":                            true,
	"semrushbot":                           true,
	"baiduspider":                          true,
	"dotbot":                               true,
	"intellij http client/goland 2024.2.3": true,
	"facebookexternalhit":                  true,
	"linkedinbot":                          true,
	"twitterbot":                           true,
	"googlebot":                            true,
	"bingbot":                              true,
	"yandexbot":                            true,
	"rogerbot":                             true,
	"embedly":                              true,
	"quora link preview":                   true,
}

const (
	geo_mm_DB = "/Users/swanhtetaungphyo/Desktop/SPI_Custom_APIGateway/backend/api/middleware/GeoLite2-Country.mmdb"
)

type Country struct {
	ISOCode string `maxminddb:"iso_code"`
	Name    string `maxminddb:"names.en"`
}

type MiddleMan struct {
	BlackSpaceIp string `yaml:"blackSpaceIp"`
}

func NewMiddleMan(blackSpaceIP string) *MiddleMan {
	return &MiddleMan{
		BlackSpaceIp: blackSpaceIP,
	}
}

func (m *MiddleMan) SetupMiddlewares(app *fiber.App) {

	app.Use(healthcheck.New())

	app.Use(m.ResponseTimeRuler)
	app.Use(m.FireWall)
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
		AllowOrigins: "http://localhost:3000, http://localhost:3001",
		AllowMethods: "GET,POST,DELETE",
		AllowHeaders: "application/json,Authorization",
	}))

	app.Use(earlydata.New(earlydata.Config{
		Error: fiber.ErrTooEarly,
	}))

	//app.Use(logger.New(logger.Config{
	//	Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
	//	Output:     os.Stdout,
	//	TimeFormat: time.RFC3339Nano,
	//	TimeZone:   "Europe/Warsaw",
	//	CustomTags: map[string]logger.LogFunc{
	//		"request_id": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
	//			return output.WriteString(c.Locals("requestid").(string))
	//		},
	//	},
	//}))

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

func (m *MiddleMan) FireWall(c *fiber.Ctx) error {
	db, err := maxminddb.Open(geo_mm_DB)
	if err != nil {
		logging.Error(err.Error())
		return err
	}

	defer func(db *maxminddb.Reader) {
		err := db.Close()
		if err != nil {
			logging.Error(err)
		}
	}(db)

	ipNet := net.ParseIP(c.IP())
	userAgent := c.Get("User-Agent")
	logging.Infof("Client User agent is %v", userAgent)
	logging.Infof("Client call from this %s", ipNet)
	var country Country
	if err := db.Lookup(ipNet, &country); err != nil {
		logging.Error(err)
		return err
	}
	if country.ISOCode == m.BlackSpaceIp {
		return c.Status(403).SendString("This website is currently forbidden in your region")
	}
	if m.isBlockUserAgent(userAgent) {
		return c.Status(403).SendString("This website cannot be accessed from your user agent")
	}
	return c.Next()
}

func (m *MiddleMan) isBlockUserAgent(agent string) bool {
	agent = strings.ToLower(agent)

	for blockedAgent := range blockedAgents {
		if strings.Contains(agent, blockedAgent) || strings.HasPrefix(agent, blockedAgent) {
			return true
		}
	}
	return false
}
