package main

import (
	"net/url"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
)

func useProxy(ctx *fiber.Ctx) error {
	if !isHttp(ctx) || ctx.Method() != fiber.MethodGet || websocket.IsWebSocketUpgrade(ctx) {
		return ctx.Next()
	}

	if !hasPort(ctx, core.ProxyPort) {
		return ctx.Redirect(core.ProxyUrl() + ctx.Path())
	}

	return ctx.Next()
}

func isHttp(ctx *fiber.Ctx) bool {
	return strings.HasPrefix(ctx.Protocol(), "http")
}

func hasPort(ctx *fiber.Ctx, port string) bool {
	u, err := url.Parse(ctx.BaseURL())
	if err != nil {
		return false
	}
	return u.Port() == port
}
