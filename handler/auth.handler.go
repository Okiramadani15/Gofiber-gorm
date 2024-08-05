package handler

import (
	"go_fiber_gorm/database"
	"go_fiber_gorm/model/entity"
	"go_fiber_gorm/model/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}
	log.Println(loginRequest)

	// // VALIDASI REQUEST
	// validate := validator.New()
	// errValidate := validate.Struct(loginRequest)
	// if errValidate != nil {
	// 	return ctx.Status(400).JSON(fiber.Map{
	// 		"message": "failed",
	// 		"error":   errValidate.Error(),
	// 	})
	// }

	var user entity.User
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong credential",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": "secret",
	})

}
