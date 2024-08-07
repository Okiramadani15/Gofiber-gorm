package route

import (
	"go_fiber_gorm/config"
	"go_fiber_gorm/handler"
	"go_fiber_gorm/middleware"
	"go_fiber_gorm/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", handler.LoginHandler)

	r.Get("/user", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)

	r.Post("/book", utils.HandleSingleFile, handler.BookHandlerCreate)

	r.Post("/gallery", utils.HandleMultipleFile, handler.PhotoHandlerCreate)

}
