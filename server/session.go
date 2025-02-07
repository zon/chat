package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
)

var store *session.Store

func initSessionStore() {
	storage := sqlite3.New()
	store = session.New(session.Config{
		Storage: storage,
	})
}

func useSession(ctx *fiber.Ctx) error {
	ok, err := checkSession(ctx)
	if err != nil {
		return err
	}

	if ctx.Method() == fiber.MethodPost && strings.HasPrefix(ctx.Path(), "/auth") {
		ok = true
	}

	if ok {
		return ctx.Next()
	} else {
		return ctx.Redirect(authUrl())
	}
}

func createSession(ctx *fiber.Ctx, userID uint) error {
	sess, err := store.Get(ctx)
	if err != nil {
		return err
	}
	sess.Set("userID", userID)
	return sess.Save()
}

func getSessionUserID(ctx *fiber.Ctx) (uint, error) {
	sess, err := store.Get(ctx)
	if err != nil {
		return 0, err
	}
	id := sess.Get("userID")
	if id == nil {
		return 0, nil
	} else {
		return id.(uint), nil
	}
}

func checkSession(ctx *fiber.Ctx) (bool, error) {
	id, err := getSessionUserID(ctx)
	return id != 0, err
}
