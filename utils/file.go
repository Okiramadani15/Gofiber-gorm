package utils

import (
	"fmt"
	"log"
	"os"

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
// }|

const DefaultpathAssetImage = "./public/covers"

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

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		log.Println("Error Read Multipart form Request, Error -", errForm)
	}
	files := form.File["photos"]

	var filenames []string
	for i, file := range files {

		var filename string
		if file != nil {
			filename = fmt.Sprintf("%d-%s", i, file.Filename)

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

		if filename != "" {
			filenames = append(filenames, filename)
		}
		ctx.Locals("filenames", filenames)
	}

	return ctx.Next()
}

func HandRemoveFile(filename string, pathFile ...string) error {
	if len(pathFile) > 0 {
		err := os.Remove(pathFile[0] + filename)
		if err != nil {
			log.Println("Failed tp remove file")
			return err
		}
	} else {
		err := os.Remove(DefaultpathAssetImage + filename)
		if err != nil {
			log.Println("Failed tp remove file")
			return err
		}
	}

	return nil
}
