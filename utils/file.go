package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const DefaultpathAssetImage = "./public/covers"

// func HandleSingleFile(ctx *fiber.Ctx) error {
// 	// HANDLE FILE
// 	file, errFile := ctx.FormFile("cover")
// 	if errFile != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Failed to get the file",
// 		})
// 	}

// 	var filename *string
// 	if file != nil {
// 		filename = &file.Filename
// 		extensionFile := filepath.Ext(*filename)
// 		newFilename := fmt.Sprintf("gambar-satu%s", extensionFile)

// 		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", newFilename))
// 		if errSaveFile != nil {
// 			log.Println("Failed to store file into public/covers directory:", errSaveFile)
// 			// return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			// 	"error": "Failed to save the file",
// 			// })
// 		}
// 	} else {
// 		log.Println("No file to be uploaded")
// 		// return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 		// 	"error": "No file uploaded",
// 		// })
// 	}

// 	ctx.Locals("filename", filename)

// 	return ctx.Next()
// }

func HandleSingleFile(ctx *fiber.Ctx) error {
	// HANDLE FILE
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get the file",
		})
	}

	var filename string
	if file != nil {
		// Ambil extension file
		extensionFile := filepath.Ext(file.Filename)
		// Buat nama file baru
		filename = fmt.Sprintf("gambar-satu%s", extensionFile)

		// Simpan file ke direktori yang diinginkan
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))
		if errSaveFile != nil {
			log.Println("Failed to store file into public/covers directory:", errSaveFile)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to save the file",
				"message": "failed",
			})
		}
	} else {
		log.Println("No file to be uploaded")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "No file uploaded",
			"message": "failed",
		})
	}

	// Simpan nama file ke dalam context locals sebagai string
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
			extensionFile := filepath.Ext(file.Filename)
			filename = fmt.Sprintf("%d-%s%s", i, "gambar", extensionFile)
			// // Ambil extension file
			// extensionFile := filepath.Ext(file.Filename)
			// // Buat nama file baru
			// filename = fmt.Sprintf("gambar-satu%s", extensionFile)

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

func HandleRemoveFile(filename string, pathFile ...string) error {
	var filePath string
	if len(pathFile) > 0 {
		filePath = pathFile[0] + filename
	} else {
		filePath = DefaultpathAssetImage + filename
	}

	err := os.Remove(filePath)
	if err != nil {
		log.Println("Failed to remove file:", filePath)
		return err
	}

	return nil
}
