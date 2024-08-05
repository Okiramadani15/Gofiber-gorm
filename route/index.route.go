package route

import (
	"go_fiber_gorm/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", handler.UserHandlerGetAll)
	r.Post("/user", handler.UserHandlerCreate)

}
