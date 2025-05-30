package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/zon/chat/core"
	"github.com/zon/chat/html"
)

type postMessageBody struct {
	Text string
}

func getMessages(c *fiber.Ctx) error {
	var messages []core.Message

	bq := c.Query("before")
	if bq != "" {
		before, err := core.ParseTime(bq)
		if err != nil {
			log.Error(err)
			return fiber.ErrBadRequest
		}
		err = core.GetMessagesBefore(before, &messages)
		if err != nil {
			return err
		}
		return render(c, html.Messages(messages, true))
	}

	aq := c.Query("after")
	if aq != "" {
		after, err := core.ParseTime(aq)
		if err != nil {
			log.Error(err)
			return fiber.ErrBadRequest
		}
		err = core.GetMessagesAfter(after, &messages)
		if err != nil {
			return err
		}
		return render(c, html.Messages(messages, false))
	}

	err := core.GetLatestMessages(&messages)
	if err != nil {
		return err
	}
	return render(c, html.Messages(messages, true))
}

func postMessage(c *fiber.Ctx) error {
	user, err := core.GetSessionUser(c)
	if err != nil {
		return err
	}

	var body postMessageBody
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}
	content := strings.TrimSpace(body.Text)

	if content == "" {
		return fiber.ErrBadRequest
	}

	content, err = core.MarkdownToHtml(content)
	if err != nil {
		return err
	}
	record, err := core.CreateMessage(*user, content)
	if err != nil {
		return err
	}

	cmp := html.OobMessage(*record)
	err = topic.RenderWrite(cmp)
	if err != nil {
		return err
	}

	return render(c, html.NewMessageForm())
}
