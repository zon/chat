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
	user, err := authUser(c)
	if err != nil {
		return err
	}
	slog.Info("auth", "authId", user.ID, "username", user.Name)
	return c.JSON(user)
}
