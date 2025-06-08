package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Nats      *url.URL
	WebSocket *url.URL
}

var config Config

func LoadConfig() error {
	nats, err := loadUrl("WURBS_NATS_URL", "WURBS_NATS_PASSWORD")
	if err != nil {
		return err
	}
	ws, err := loadUrl("WURBS_WS_URL", "WURBS_WS_PASSWORD")
	if err != nil {
		return err
	}
	config = Config{
		Nats:      nats,
		WebSocket: ws,
	}
	return nil
}

func loadUrl(urlEnv string, passwordEnv string) (*url.URL, error) {
	u := &url.URL{}
	v := os.Getenv(urlEnv)
	if v == "" {
		return u, fmt.Errorf("%s env var is missing", urlEnv)
	}
	u, err := url.Parse(v)
	if err != nil {
		return u, err
	}
	v = os.Getenv(passwordEnv)
	if v == "" {
		return u, fmt.Errorf("%s env var is missing", passwordEnv)
	}
	u.User = url.UserPassword(u.User.Username(), v)
	return u, nil
}

func getWebsocket(c *fiber.Ctx) error {
	return c.JSON(config.WebSocket.String())
}
