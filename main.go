package main

import (
	"online-store-application/database"
	"online-store-application/database/migration"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDB()
	migration.InitMigration()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
