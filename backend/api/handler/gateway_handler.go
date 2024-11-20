package handler

import (
	"github.com/SwanHtetAungPhyo/api/models"
	"github.com/gofiber/fiber/v2"
)

type GateWayHandler struct {
}

func (g *GateWayHandler) Services(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(models.Configuration())
}
