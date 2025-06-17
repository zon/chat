package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Nats      *url.URL
	WebSocket WebSocketConfig
}

type WebSocketConfig struct {
	Host     string
	User     string
	Password string
}

var config Config

func LoadConfig() error {
	nats, err := loadUrl("WURBS_NATS_URL", "WURBS_NATS_PASSWORD")
	if err != nil {
		return err
	}

	host, err := requireEnv("WURBS_WS_HOST")
	if err != nil {
		return err
	}
	user, err := requireEnv("WURBS_WS_USER")
	if err != nil {
		return err
	}
	password, err := requireEnv("WURBS_WS_PASSWORD")
	if err != nil {
		return err
	}

	config = Config{
		Nats: nats,
		WebSocket: WebSocketConfig{
			Host:     host,
			User:     user,
			Password: password,
		},
	}
	return nil
}

func loadUrl(urlEnv string, passwordEnv string) (*url.URL, error) {
	u := &url.URL{}
	v, err := requireEnv(urlEnv)
	if err != nil {
		return u, err
	}
	u, err = url.Parse(v)
	if err != nil {
		return u, err
	}
	v, err = requireEnv(passwordEnv)
	if err != nil {
		return u, err
	}
	u.User = url.UserPassword(u.User.Username(), v)
	return u, nil
}

func requireEnv(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return value, fmt.Errorf("%s env var is missing", name)
	}
	return value, nil
}

func getWebsocket(c *fiber.Ctx) error {
	return c.JSON(config.WebSocket)
}
