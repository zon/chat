package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

const proxyPort string = "7331"

func useProxy(ctx *fiber.Ctx) error {
	if !isHttp(ctx) || ctx.Method() != fiber.MethodGet || websocket.IsWebSocketUpgrade(ctx) {
		return ctx.Next()
	}

	if !hasPort(ctx, proxyPort) {

		fmt.Println("redirect", proxyUrl())

		return ctx.Redirect(proxyUrl() + ctx.Path())
	}

	return ctx.Next()
}

func isHttp(ctx *fiber.Ctx) bool {
	return strings.HasPrefix(ctx.Protocol(), "http")
}

func hasPort(ctx *fiber.Ctx, port string) bool {
	return strings.HasSuffix(ctx.BaseURL(), ":"+port)
}

func proxyUrl() string {
	return protocol + "://" + host() + ":" + proxyPort
}
