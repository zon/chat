package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
	"gorm.io/gorm"
)

type userPostBody struct {
	Name string
}

func getUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	user, err := core.GetUser(uint(id))
	if err != nil {
		return err
	}
	if user == nil {
		return fiber.ErrNotFound
	}
	return c.JSON(user)
}

func putUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := authUser(c)
	if err != nil {
		return err
	}
	if user.ID != uint(id) {
		return fiber.ErrForbidden
	}

	var body userPostBody
	err = c.BodyParser(&body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user.Name = body.Name
	user.Ready = true
	err = user.Save()
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return respondBad(c, Duplicate, "Duplicate user name")
	}
	if err != nil {
		return err
	}

	return c.JSON(user)
}
