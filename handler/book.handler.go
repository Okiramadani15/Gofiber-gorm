package handler

import (
	"go_fiber_gorm/database"
	"go_fiber_gorm/model/entity"
	"go_fiber_gorm/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// func BookHandlerCreate(ctx *fiber.Ctx) error {
// 	book := new(request.BookCreateRequest)
// 	if err := ctx.BodyParser(book); err != nil {
// 		return err
// 	}

// 	// VALIDASI REQUEST

// 	validate := validator.New()
// 	errValidate := validate.Struct(book)
// 	if errValidate != nil {
// 		return ctx.Status(400).JSON(fiber.Map{
// 			"message": "failed",
// 			"error":   errValidate.Error(),
// 		})
// 	}

// 	// // HANDLE FILE
// 	filename := ctx.Locals("filename").(string)

// 	newBook := entity.Book{
// 		Title:  book.Title,
// 		Author: book.Author,
// 		Cover:  filename,
// 	}

// 	errCreateBook := database.DB.Create(&newBook).Error
// 	if errCreateBook != nil {
// 		return ctx.Status(500).JSON(fiber.Map{
// 			"message": "failed to store data",
// 		})
// 	}

// 	return ctx.JSON(fiber.Map{
// 		"message": "success",
// 		"data":    newBook,
// 	})
// }

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

	// VALIDATION REQUIRED IMAGE
	filenameLocal := ctx.Locals("filename")
	if filenameLocal == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image cover is required",
		})
	}

	filename, ok := filenameLocal.(string)
	if !ok {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed",
			"error":   "filename is not a string",
		})
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
