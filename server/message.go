package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
	"github.com/zon/chat/html"
)

type postMessageBody struct {
	Text string
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
	content := markdownToHtml(body.Text)

	err = render(c, html.NewMessage())
	if err != nil {
		return err
	}

	now := time.Now()

	return render(c, html.OobMessage(user.Name, now, content))
}
