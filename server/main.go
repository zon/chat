package main

import (
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
	"github.com/zon/chat/net"
)

var topic *net.Topic

func main() {
	proxy := os.Getenv("PROXY") != "false"

	err := core.InitDB(true)
	if err != nil {
		log.Fatal(err)
	}
	core.InitSessionStore()
	LoadAuthTokenSecret()

	topic = net.MakeTopic()

	app := fiber.New()

	if proxy {
		app.Use(useProxy)
	}
	app.Use(useSession)
	app.Get("/", getIndex)
	app.Get("/messages", getMessages)
	app.Post("/", postMessage)
	app.Post("/auth", postAuth)
	app.Get("/user/:id", getUser)
	app.Post("/user/:id", postUser)
	
	app.Use("/ws", useWebsocket)
	app.Get("/ws/:id", websocket.New(handleWebocket))

	app.Listen(":"+ port)
}
