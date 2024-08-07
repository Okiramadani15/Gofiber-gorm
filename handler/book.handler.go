package handler

import (
	"fmt"
	"go_fiber_gorm/database"
	"go_fiber_gorm/model/entity"
	"go_fiber_gorm/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)
	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	// VALIDASI REQUEST

	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// HANDLE FILE
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error File =", errFile)
	}

	var filename string
	if file != nil {
		filename = file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))
		if errSaveFile != nil {
			log.Println("Fail to store file into public/covers directory.")
		}
	} else {
		log.Println("nothing file to be uploading")
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filename,
	}

	errCreateBook := database.DB.Create(&newBook).Error
	if errCreateBook != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newBook,
	})
}
