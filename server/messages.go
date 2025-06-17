package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
)

func getMessages(c *fiber.Ctx) error {
	var messages []core.Message

	bq := c.Query("before")
	if bq != "" {
		before, err := core.ParseTime(bq)
		if err != nil {
			return fiber.ErrBadRequest
		}
		err = core.GetMessagesBefore(before, &messages)
		if err != nil {
			return err
		}
		return c.JSON(messages)
	}

	aq := c.Query("after")
	if aq != "" {
		after, err := core.ParseTime(aq)
		if err != nil {
			return fiber.ErrBadRequest
		}
		err = core.GetMessagesAfter(after, &messages)
		if err != nil {
			return err
		}
		return c.JSON(messages)
	}

	err := core.GetLatestMessages(&messages)
	if err != nil {
		return err
	}
	return c.JSON(messages)
}

func postMessage(c *fiber.Ctx) error {
	user, err := authUser(c)
	if err != nil {
		return err
	}

	var body string
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}
	content := strings.TrimSpace(body)

	if content == "" {
		return fiber.ErrBadRequest
	}

	content, err = core.MarkdownToHtml(content)
	if err != nil {
		return err
	}
	record, err := core.CreateMessage(user.ID, content)
	if err != nil {
		return err
	}

	err = Publish("messages", record)
	if err != nil {
		return err
	}

	return c.JSON(record)
}
