package handler

import (
	"fmt"
	"go_fiber_gorm/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PhotoHandlerCreate(ctx *fiber.Ctx) error {
	photo := new(request.PhotoCreateRequest)
	if err := ctx.BodyParser(photo); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(photo)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// VALIDATION REQUIRED IMAGE
	var filenameString string

	filenames := ctx.Locals("filename")
	if filenames == nil {
		return ctx.Status(200).JSON(fiber.Map{
			"message": "Success added photos",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filenames)
	}

	log.Println(filenameString)

	// filenames, ok := filenames.(string)
	// if !ok {
	// 	return ctx.Status(500).JSON(fiber.Map{
	// 		"message": "failed",
	// 		"error":   "filename is not a string",
	// 	})
	// }

	// newPhoto := entity.Photo{
	// 	Image:      filename,
	// 	CategoryID: 1,
	// }

	// errCreatePhoto := database.DB.Create(&newPhoto).Error
	// if errCreatePhoto != nil {
	// 	log.Println("ada file gagal")
	// 	return ctx.Status(500).JSON(fiber.Map{
	// 		"message": "failed to store data",
	// 	})
	// }

	return ctx.JSON(fiber.Map{
		"message": "success",
		// "data":    photo,
	})
}
