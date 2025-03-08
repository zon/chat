package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/zon/chat/core"
)

func useSession(ctx *fiber.Ctx) error {
	session, err := core.GetSession(ctx)
	if err != nil {
		return err
	}

	log.Debugf("%s %s session %s %d", ctx.Method(), ctx.Path(), session.ID(), session.UserID)

	if !session.HasAuth() {
		if isPost(ctx) && isResource(ctx, core.AuthPath) {
			return ctx.Next()
		} else {
			log.Debugf("%s %s %s", ctx.Method(), ctx.Path(), "unauthenticated")
			return ctx.Redirect(core.AuthUrl())
		}
	}

	if !session.Ready {
		if isResource(ctx, core.UserPath) {
			return ctx.Next()
		} else {
			log.Debugf("%s %s %s", ctx.Method(), ctx.Path(), "not ready")
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
