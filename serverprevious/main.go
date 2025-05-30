package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
	"github.com/zon/chat/net"
)

var cli struct {
	Proxy bool `help:"Redirect to templ watch proxy"`
}

var topic *net.Topic

func main() {
	ctx := kong.Parse(&cli)
	ctx.FatalIfErrorf(ctx.Error)

	core.InitConfig()
	err := core.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	core.InitSessionStore()
	LoadAuthTokenSecret()

	topic = net.MakeTopic()

	app := fiber.New()

	app.Use(useSession)
	if cli.Proxy {
		app.Use(useProxy)
	}
	app.Get("/", getIndex)
	app.Get("/messages", getMessages)
	app.Post("/", postMessage)
	app.Post("/auth", postAuth)
	app.Get("/user/:id", getUser)
	app.Post("/user/:id", postUser)

	app.Use("/ws", useWebsocket)
	app.Get("/ws/:id", websocket.New(handleWebocket))

	app.Listen(":" + core.Port())
}
