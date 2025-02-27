package net

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/contrib/websocket"
)

var autoId int = 0

type Client struct {
	Id     int
	UserID uint
	conn   *websocket.Conn
}

type Message struct {
	Text    string
	Headers map[string]string
}

func MakeClient(conn *websocket.Conn, userIDKey string) *Client {
	autoId += 1
	userID := conn.Locals(userIDKey).(uint)
	return &Client{
		Id:     autoId,
		UserID: userID,
		conn:   conn,
	}
}

func (c *Client) ReadMessage(msg *Message) error {
	t, d, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}
	if t != websocket.TextMessage {
		return fmt.Errorf("non text message type %d", t)
	}
	err = json.Unmarshal(d, &msg)
	if err != nil {
		return err
	}
	msg.Text = strings.TrimSpace(msg.Text)

	log.Printf("%d received %s", c.Id, msg)

	return nil
}

func (c *Client) WriteText(data []byte) error {

	log.Printf("%d sent %s", c.Id, string(data))

	return c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *Client) RenderWrite(cmp templ.Component) error {
	writer, err := c.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	err = cmp.Render(context.Background(), writer)
	if err != nil {
		return err
	}
	return writer.Close()
}

func (c *Client) Close() error {
	return c.conn.Conn.Close()
}

func (m *Message) IsEmpty() bool {
	return m.Text == ""
}