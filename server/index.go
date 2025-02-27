package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
	"github.com/zon/chat/html"
)

func getIndex(c *fiber.Ctx) error {
	user, err := core.GetSessionUser(c)
	if err != nil {
		return err
	}

	var messages []core.Message
	err = core.GetLatestMessages(&messages)
	if err != nil {
		return err
	}

	alert := getAlert(c)
	return renderBody(c, html.Index(user, messages, wsUrl(), alert))
}

func getAlert(c *fiber.Ctx) string {
	alert := c.Locals(alertKey)
	if alert == nil {
		return ""
	}
	return alert.(string)
}

func redirectToIndex(c *fiber.Ctx, alert string) error {
	pushUrl := "/"
	if isHxRequest(c) {
		c.Set("HX-Push-Url", pushUrl)
		c.Locals(alertKey, alert)
		return getIndex(c)
	} else {
		c.Redirect(pushUrl)
		return nil
	}
}
