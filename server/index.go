package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
	"github.com/zon/chat/html"
)

func getIndex(c *fiber.Ctx) error {
	session, err := core.GetSession(c)
	if err != nil {
		return err
	}
	user, err := core.GetUser(session.UserID)
	if err != nil {
		return err
	}
	return renderIndex(c, user, "")
}

func renderIndex(c *fiber.Ctx, user *core.User, alert string) error {
	return renderBody(c, html.Index(user, alert))
}

func redirectToIndex(ctx *fiber.Ctx, user *core.User, alert string) error {
	pushUrl := "/"
	if isHxRequest(ctx) {
		ctx.Set("HX-Push-Url", pushUrl)
		return renderIndex(ctx, user, alert)
	} else {
		ctx.Redirect(pushUrl)
		return nil
	}
}
