package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
	"github.com/zon/chat/net"
)

var topic *net.Topic

func main() {
	err := core.InitDB(true)
	if err != nil {
		log.Fatal(err)
	}
	core.InitSessionStore()
	LoadAuthTokenSecret()

	topic = net.MakeTopic()

	app := fiber.New()

	app.Use(useProxy)
	app.Use(useSession)
	app.Get("/", getIndex)
	app.Get("/messages", getMessages)
	app.Post("/auth", postAuth)
	app.Get("/user/:id", getUser)
	app.Post("/user/:id", postUser)
	
	app.Use("/ws", useWebsocket)
	app.Get("/ws/:id", websocket.New(handleWebocket))

	app.Listen(":"+ port)
}
