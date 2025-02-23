package main

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/html"
)

func isHxRequest(ctx *fiber.Ctx) bool {
	return ctx.Get("Hx-Request", "false") == "true"
}

func render(ctx *fiber.Ctx, cmp templ.Component) error {
	ctx.Type("html")
	return cmp.Render(ctx.Context(), ctx)
}

func renderLayout(ctx *fiber.Ctx, cmp templ.Component) error {
	return renderBody(ctx, html.Layout(cmp))
}

func renderBody(ctx *fiber.Ctx, cmp templ.Component) error {
	if isHxRequest(ctx) {
		return render(ctx, cmp)
	} else {
		return render(ctx, html.Doc(cmp))
	}
}

func renderError(ctx *fiber.Ctx, message string) error {
	ctx.Set("HX-Reswap", "none")
	return render(ctx, html.OobError(message))
}
