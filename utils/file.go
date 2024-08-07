package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// func HandleSingleFile(ctx *fiber.Ctx) error {
// 	// HANDLE FILE
// 	file, errFile := ctx.FormFile("cover")
// 	if errFile != nil {
// 		log.Println("Error File =", errFile)
// 	}

// 	var filename *string
// 	if file != nil {
// 		filename = &file.Filename

// 		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", *filename))
// 		if errSaveFile != nil {
// 			log.Println("Fail to store file into public/covers directory.")
// 		}
// 	} else {
// 		log.Println("nothing file to be uploading")
// 	}

// 	ctx.Locals("filename", filename)

// 	return ctx.Next()
// }

func HandleSingleFile(ctx *fiber.Ctx) error {
	// HANDLE FILE
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error File =", errFile)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get the file",
		})
	}

	var filename string
	if file != nil {
		filename = file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))
		if errSaveFile != nil {
			log.Println("Failed to store file into public/covers directory:", errSaveFile)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save the file",
			})
		}
	} else {
		log.Println("No file to be uploaded")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	ctx.Locals("filename", filename)

	return ctx.Next()
}
