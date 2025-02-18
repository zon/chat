package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

const proxyPort string = "7331"

func useProxy(ctx *fiber.Ctx) error {
	if ctx.Method() != fiber.MethodGet {
		return ctx.Next()
	}
	
	if !strings.HasSuffix(ctx.BaseURL(), ":"+ proxyPort) {
		return ctx.Redirect(proxyUrl() + ctx.Path())
	}

	return ctx.Next()
}

func proxyUrl() string {
	return protocol + "://" + host + ":" + proxyPort
}