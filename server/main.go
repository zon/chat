package main

import (
	"log/slog"
	"os"

	"github.com/alecthomas/kong"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/zon/chat/core"
	"github.com/zon/gonf"
)

var cli struct {
	Key       string `arg:"" type:"existingfile" help:"Path to Zitadel API private key json file"`
	Subdomain string `help:"Zitadel application subdomain" default:"wurbs-2d2isd"`
	Port      string `help:"Port to host on" default:"8080"`
}

func main() {
	ktx := kong.Parse(&cli)
	ktx.FatalIfErrorf(ktx.Error)

	err := LoadConfig()
	if err != nil {
		slog.Error("config", "error", err)
		os.Exit(1)
	}

	err = core.InitDB()
	if err != nil {
		slog.Error("db", "error", err)
		os.Exit(1)
	}

	err = gonf.InitAuthMiddleware(cli.Subdomain, cli.Key)
	if err != nil {
		slog.Error("zitadel auth middleware could not initialize", "error", err)
		os.Exit(1)
	}

	err = Connect()
	if err != nil {
		slog.Error("nats connection failed", "error", err)
		os.Exit(1)
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("ok")
	})

	app.Use(gonf.AuthMiddleware)

	app.Get("/auth", gonf.GetAuth)
	app.Get("/websocket", getWebsocket)
	app.Get("/messages", getMessages)
	app.Post("/messages", postMessage)
	app.Get("/users", getUsers)
	app.Get("/users/:id", getUser)
	app.Put("/users/:id", putUser)

	err = app.Listen(":" + cli.Port)
	if err != nil {
		slog.Error("Listen failed", "error", err)
	}
}
