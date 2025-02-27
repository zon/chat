package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
)

func useSession(ctx *fiber.Ctx) error {
	session, err := core.GetSession(ctx)
	if err != nil {
		return err
	}

	if !session.HasAuth() {
		if isPost(ctx) && isResource(ctx, authPath) {
			return ctx.Next()
		} else {
			return ctx.Redirect(authUrl())
		}
	}

	if !session.Ready {
		if isResource(ctx, core.UserPath) {
			return ctx.Next()
		} else {
			return ctx.Redirect(core.UserUrl(session.UserID))
		}
	}

	ctx.Locals(userIDKey, session.UserID)

	return ctx.Next()
}

func isPost(ctx *fiber.Ctx) bool {
	return ctx.Method() == fiber.MethodPost
}

func isResource(ctx *fiber.Ctx, path string) bool {
	return strings.HasPrefix(ctx.Path(), path)
}
