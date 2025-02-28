package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/net"
)

func wsUrl() string {
	return url() + "/ws/123?v=1.0"
}

func useWebsocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func handleWebocket(c *websocket.Conn) {
	client := net.MakeClient(c, userIDKey)
	topic.Join(client)
	log.Println(client.Id, "joined")

	var msg net.Message
	var err error
	for {
		err = client.ReadMessage(&msg)
		if err != nil {
			log.Println(client.Id, "read error:", err)
			break
		}
		if msg.IsEmpty() {
			continue
		}
	}
	topic.Leave(client)
	err = client.Close()
	if err != nil {
		log.Println(client.Id, "close error:", err)
		return
	}
	log.Println(client.Id, "closed")
}
