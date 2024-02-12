package route

import (
	"online-store-application/handler"
	"online-store-application/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(r *fiber.App) {
	r.Post("/users/register", handler.UserHandlerRegister)
	r.Post("/users/login", handler.LoginHandler)
	r.Post("users/logout", middleware.Auth, handler.LogoutHandler)
}
