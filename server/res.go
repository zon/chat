package main

import "github.com/gofiber/fiber/v2"

const Duplicate string = "duplicate"

type BadRequest struct {
	Code    string
	Message string
}

func respondBad(c *fiber.Ctx, code string, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(BadRequest{
		Code:    code,
		Message: message,
	})
}
