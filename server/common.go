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
	layout := html.Layout(cmp)
	if isHxRequest(ctx) {
		return render(ctx, layout)
	} else {
		return render(ctx, html.Doc(layout))
	}
}

func renderError(ctx *fiber.Ctx, message string) error {
	ctx.Set("HX-Reswap", "none")
	return render(ctx, html.OobError(message))
}

func compact(list []string) []string {
	result := []string{}
	for _, item := range list {
		if item != "" {
			result = append(result, item)
		}
	}
	return result
}

func head(list []string) string {
	if len(list) < 1 {
		return ""
	}
	return list[0]
}

func tail(list []string) string {
	c := len(list)
	if c < 1 {
		return ""
	}
	return list[c-1]
}
