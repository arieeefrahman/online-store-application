package middleware

import (
	"errors"
	"online-store-application/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		err := errors.New("missing authorization header")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "failed to authenticate",
			"error":   err.Error(),
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		err := errors.New("invalid authorization header format")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "failed to authenticate",
			"error":   err.Error(),
		})	
	}
	token := parts[1]

	_, err := util.VerifyToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "failed to authenticate",
			"error":   err.Error(),
		})
	}

	return ctx.Next()
}
