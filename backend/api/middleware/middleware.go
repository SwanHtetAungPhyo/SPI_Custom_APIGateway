package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type MiddleMan struct {
}

func (m *MiddleMan) ResponseTimeRuler(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)
	key := fmt.Sprintf("Response time from %s", c.OriginalURL())
	c.Set(key, strconv.FormatInt(int64(duration), 10))
	return err
}
