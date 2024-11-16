package handler

import (
	"github.com/SwanHtetAungPhyo/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {

	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if request.Username == "user" && request.Password == "password" {

		token, err := middleware.GenerateJWT(request.Username)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Could not generate token")
		}
		return ctx.JSON(fiber.Map{
			"token": token,
		})
	}

	return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
}
