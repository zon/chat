package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/alecthomas/kong"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

var cli struct {
	Key       string `arg:"" type:"existingfile" help:"Path to Zitadel API private key json file"`
	Subdomain string `help:"Zitadel application subdomain" default:"wurbs-2d2isd"`
	Port      string `help:"Port to host on" default:"8080"`
}

type User struct {
	ID   string
	Name string
}

func main() {
	ktx := kong.Parse(&cli)
	ktx.FatalIfErrorf(ktx.Error)

	ctx := context.Background()

	domain := fmt.Sprintf("%s.us1.zitadel.cloud", cli.Subdomain)
	authZ, err := authorization.New(ctx, zitadel.New(domain), oauth.DefaultAuthorization(cli.Key))
	if err != nil {
		slog.Error("zitadel sdk could not initialize", "error", err)
		os.Exit(1)
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("ok")
	})

	mw := middleware.New(authZ)
	app.Use(adaptor.HTTPMiddleware(mw.RequireAuthorization()))

	app.Get("/session", func(c *fiber.Ctx) error {
		aCtx := mw.Context(c.Context())
		slog.Info("user accessed task list", "id", aCtx.UserID(), "username", aCtx.Username)
		user := User { ID: aCtx.UserID(), Name: aCtx.Username }
		return c.JSON(user)
	})

	err = app.Listen(":" + cli.Port)
	if err != nil {
		slog.Error("Listen failed", "error", err)
	}
}
