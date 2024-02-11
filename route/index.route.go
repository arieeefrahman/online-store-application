package route

import (
	"online-store-application/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(r *fiber.App) {
	r.Post("/users/register", handler.UserHandlerRegister)
	r.Post("/users/login", handler.LoginHandler)
}
