package main

import (
	"go_fiber_gorm/database"
	"go_fiber_gorm/database/migration"
	"go_fiber_gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INITIAL DATABASE
	database.DatabaseInit()

	//RUN MIGRATION
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":8080")
}
