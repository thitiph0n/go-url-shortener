package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thitiph0n/go-url-shortener/errs"
)

func HandleError(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	message := "Something went wrong :("

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	switch e := err.(type) {
	case errs.AppError:
		code = e.Code
		message = e.Message

	case *fiber.Error:
		code = e.Code

	case error:
		message = e.Error()
	}

	return c.Status(code).JSON(fiber.Map{"error": message})

}
