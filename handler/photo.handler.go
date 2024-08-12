package handler

import (
	"go_fiber_gorm/database"
	"go_fiber_gorm/model/entity"
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

	// CEK VALIDASI CategoryID
	var categoryCount int64
	database.DB.Model(&entity.Category{}).Where("id = ?", photo.Categoryid).Count(&categoryCount)
	if categoryCount == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid CategoryID",
		})
	}

	// VALIDATION REQUIRED IMAGE
	filenames := ctx.Locals("filenames")
	log.Println("filename =", filenames)
	if filenames == nil {
		return ctx.Status(442).JSON(fiber.Map{
			"message": "Image Cover id Required",
		})
	} else {
		filenamesData := filenames.([]string)
		for _, filename := range filenamesData {
			newPhoto := entity.Photo{
				Image:      filename,
				CategoryID: photo.Categoryid,
			}

			errCreatePhoto := database.DB.Create(&newPhoto).Error
			if errCreatePhoto != nil {
				log.Println("some data not saved properly")
				return ctx.Status(500).JSON(fiber.Map{
					"message": "failed to store data",
				})
			}
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

// func PhotoHandlerCreate(ctx *fiber.Ctx) error {
// 	photo := new(request.PhotoCreateRequest)
// 	if err := ctx.BodyParser(photo); err != nil {
// 		return err
// 	}

// 	// VALIDASI REQUEST
// 	validate := validator.New()
// 	errValidate := validate.Struct(photo)
// 	if errValidate != nil {
// 		return ctx.Status(400).JSON(fiber.Map{
// 			"message": "failed",
// 			"error":   errValidate.Error(),
// 		})
// 	}

// 	// VALIDATION REQUIRED IMAGE
// 	//var filenameString string

// 	filenames := ctx.Locals("filenames")
// 	log.Println("filename =", filenames)
// 	if filenames == nil {
// 		return ctx.Status(442).JSON(fiber.Map{
// 			"message": "Image Cover id Required",
// 		})
// 	} else {
// 		// filenameString = fmt.Sprintf("%v", filenames)
// 		filenamesData := filenames.([]string)
// 		for _, filename := range filenamesData {

// 			newPhoto := entity.Photo{
// 				Image:      filename,
// 				CategoryID: photo.Categoryid,
// 			}

// 			errCreatePhoto := database.DB.Create(&newPhoto).Error
// 			if errCreatePhoto != nil {
// 				log.Println("somed data not saved properly")
// 				return ctx.Status(500).JSON(fiber.Map{
// 					"message": "failed to store data",
// 				})
// 			}
// 		}
// 	}

// 	// log.Println("filenames :: ", filenameString)

// 	return ctx.JSON(fiber.Map{
// 		"message": "success",
// 		// "data":    photo,
// 	})
// }
