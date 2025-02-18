package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
)

func main() {
	err := core.InitDB(true)
	if err != nil {
		log.Fatal(err)
	}
	core.InitSessionStore()
	LoadAuthTokenSecret()

	app := fiber.New()

	app.Use(useSession)
	app.Get("/", getIndex)
	app.Post("/auth", postAuth)
	app.Get("/user/:id", getUser)
	app.Post("/user/:id", getUser)

	app.Listen(":8080")
}
