package handler

import (
	"go_fiber_gorm/database"
	"go_fiber_gorm/model/entity"
	"go_fiber_gorm/model/request"
	"go_fiber_gorm/utils"
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

func PhotoHandlerDelete(ctx *fiber.Ctx) error {
	photoId := ctx.Params("id")
	var photo entity.Photo

	// CHECK AVAILABLE PHOTO
	err := database.DB.First(&photo, "id = ?", photoId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Photo not found",
		})
	}

	// Pastikan bahwa photo.Image tidak nil atau kosong
	if photo.Image == "" {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "No image associated with this photo record",
		})
	}

	// HANDLE REMOVE FILE (hapus file cover juga)
	errDeleteFile := utils.HandleRemoveFile(photo.Image, "./public/covers/")
	if errDeleteFile != nil {
		log.Println("fail to delete the file:", errDeleteFile)
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to delete the associated image file",
		})
	}

	// Delete record from database
	errDelete := database.DB.Debug().Delete(&photo).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "photo was deleted successfully",
	})
}
