package main

import (
	"online-store-application/database"
	"online-store-application/database/migration"
	"online-store-application/redis"
	"online-store-application/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	redis.InitRedis()
	database.InitDB()
	migration.InitMigration()
	app := fiber.New()
	route.InitRoute(app)
	app.Listen(":8080")
}
