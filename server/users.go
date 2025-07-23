package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
	"github.com/zon/gonf"
	"gorm.io/gorm"
)

type userPostBody struct {
	Name string
}

func getUsers(c *fiber.Ctx) error {
	aq := c.Query("after")
	if aq == "" {
		return fiber.ErrBadRequest
	}
	after, err := core.ParseTime(aq)
	if err != nil {
		return fiber.ErrBadRequest
	}

	var users []gonf.User
	err = gonf.GetUsersAfter(after, &users)
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	user, err := gonf.GetUser(uint(id))
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

	user, err := gonf.AuthUser(c)
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
		return gonf.RespondBad(c, gonf.Duplicate, "Duplicate user name")
	}
	if err != nil {
		return err
	}

	err = gonf.Publish("users", user)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
