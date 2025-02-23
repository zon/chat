package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/html"
)

func postMessage(c *fiber.Ctx) error {
	fmt.Println("body", string(c.Body()))
	return render(c, html.NewMessage())
}