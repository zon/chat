package main

import (
	"github.com/zon/chat/core"
	"github.com/zon/chat/html"
	"github.com/zon/chat/net"
)

func handleMessage(c *net.Client, msg *net.Message) error {
	user, err := core.GetUser(c.UserID)
	if err != nil {
		return err
	}

	content := markdownToHtml(msg.Text)
	record, err := core.CreateMessage(c.UserID, content)
	if err != nil {
		return err
	}

	cmp := html.NewMessageForm()
	err = topic.RenderWrite(cmp)
	if err != nil {
		return err
	}

	cmp = html.OobMessage(record, user)
	return topic.RenderWrite(cmp)
}
