package route

import (
	"online-store-application/handler"
	"online-store-application/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRoute(r *fiber.App) {
	userRoutes := r.Group("/users")
	userRoutes.Post("/register", handler.UserHandlerRegister)
	userRoutes.Post("/login", handler.LoginHandler)
	userRoutes.Post("/logout", middleware.Auth, handler.LogoutHandler)

	productRoutes := r.Group("/products", middleware.Auth)
	productRoutes.Post("/", handler.ProductCreateHandler)
	productRoutes.Get("/", handler.ProductGetAllHandler)
	productRoutes.Get("/category/:category_id", handler.ProductGetByCategoryHandler)
	productRoutes.Get("/:product_id", handler.ProductGetByIdHandler)
	productRoutes.Delete("/:product_id", handler.ProductDeleteHandler)

	cartItemRoutes := r.Group("/cart-items", middleware.Auth)
	cartItemRoutes.Post("/", handler.CartItemCreateHandler)
	cartItemRoutes.Put("/:cart_item_id", handler.CartItemUpdateHandler)
	cartItemRoutes.Delete("/:cart_item_id", handler.CartItemDeleteHandler)
	cartItemRoutes.Get("/user", handler.CartItemGetByUserIdHandler)
}
