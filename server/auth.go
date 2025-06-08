package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

type Auth struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var authMiddleware *middleware.Interceptor[*oauth.IntrospectionContext]

func initAuthMiddleware(subdomain, key string) error {
	ctx := context.Background()

	domain := fmt.Sprintf("%s.us1.zitadel.cloud", subdomain)
	authZ, err := authorization.New(ctx, zitadel.New(domain), oauth.DefaultAuthorization(key))
	if err != nil {
		return err
	}

	authMiddleware = middleware.New(authZ)
	return nil
}

func authContext(c *fiber.Ctx) *oauth.IntrospectionContext {
	return authMiddleware.Context(c.Context())
}

func authUser(c *fiber.Ctx) (*core.User, error) {
	return core.GetUserByAuthID(authContext(c).UserID())
}

func getAuth(c *fiber.Ctx) error {
	ac := authContext(c)
	slog.Info("session", "id", ac.UserID(), "username", ac.Username)
	auth := Auth{ID: ac.UserID(), Name: ac.Username}
	return c.JSON(auth)
}
