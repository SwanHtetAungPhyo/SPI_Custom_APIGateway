package handler

//
//import (
//	"github.com/SwanHtetAungPhyo/api/middleware"
//	"github.com/SwanHtetAungPhyo/api/models"
//	"github.com/gofiber/fiber/v2"
//)
//
//func LoginHandler(ctx *fiber.Ctx, config *models.GatewayConfig) error {
//
//	var request struct {
//		Username string `json:"username"`
//		Password string `json:"password"`
//	}
//
//	if err := ctx.BodyParser(&request); err != nil {
//		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
//	}
//
//	if request.Username == "user" && request.Password == "password" {
//
//		token, err := middleware.GenerateJWT(request.Username, config, "ACCESS")
//		refreshToken, err := middleware.GenerateJWT(request.Username, config, "ACCESS")
//		if err != nil {
//			return fiber.NewError(fiber.StatusInternalServerError, "Could not generate token")
//		}
//		return ctx.JSON(fiber.Map{
//			"ACCESS": token,
//			"REFRESH": refreshToken,
//		})
//	}
//
//	return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
//}
