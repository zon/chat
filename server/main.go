package main

import "github.com/gofiber/fiber/v2"

func main() {
	initSessionStore()
	
	app := fiber.New()

	app.Use(useSession)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	app.Listen(":8080")
}
