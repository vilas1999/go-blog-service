package security

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {

	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized: No Authorization header provided",
		})
	}

	user, err := AuthenticateJWT(authHeader)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized: Invalid JWT token" + err.Error(),
		})
	}

	ctx.Locals("user", user)
	return ctx.Next()
}
