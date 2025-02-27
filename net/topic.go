package net

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
)

type Topic struct {
	clients map[*Client]bool
}

func MakeTopic() *Topic {
	return &Topic{
		clients: map[*Client]bool{},
	}
}

func (t *Topic) Join(client *Client) {
	t.clients[client] = true
}

func (t *Topic) Leave(client *Client) {
	delete(t.clients, client)
}

func (t *Topic) RenderWrite(cmp templ.Component) error {
	var buf bytes.Buffer
	writer := io.Writer(&buf)
	err := cmp.Render(context.Background(), writer)
	if err != nil {
		return err
	}
	data := buf.Bytes()
	for client := range t.clients {
		err = client.WriteText(data)
		if err != nil {
			return err
		}
	}
	return nil
}
