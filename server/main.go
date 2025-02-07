package main

import "github.com/gofiber/fiber/v2"

func main() {
	initSessionStore()
	LoadAuthTokenSecret()
	
	app := fiber.New()

	app.Use(useSession)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})
	app.Post("/auth", postAuth)

	app.Listen(":8080")
}
