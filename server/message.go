package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/zon/chat/core"
	"github.com/zon/chat/html"
	"github.com/zon/chat/net"
)

func getMessages(c *fiber.Ctx) error {
	var messages []core.Message

	bq := c.Query("before")
	if bq != "" {
		before, err := core.ParseTime(bq)
		if err != nil {
			log.Error(err)
			return fiber.ErrBadRequest
		}
		err = core.GetMessagesBefore(before, &messages)
		if err != nil {
			return err
		}
		return render(c, html.Messages(messages, true))
	}

	aq := c.Query("after")
	if aq != "" {
		after, err := core.ParseTime(aq)
		if err != nil {
			log.Error(err)
			return fiber.ErrBadRequest
		}
		err = core.GetMessagesAfter(after, &messages)
		if err != nil {
			return err
		}
		return render(c, html.Messages(messages, false))
	}

	err := core.GetLatestMessages(&messages)
	if err != nil {
		return err
	}
	return render(c, html.Messages(messages, true))
}

func handleMessage(c *net.Client, msg *net.Message) error {
	user, err := core.GetUser(c.UserID)
	if err != nil {
		return err
	}

	content := markdownToHtml(msg.Text)
	record, err := core.CreateMessage(*user, content)
	if err != nil {
		return err
	}

	cmp := html.NewMessageForm()
	err = topic.RenderWrite(cmp)
	if err != nil {
		return err
	}

	cmp = html.OobMessage(*record)
	return topic.RenderWrite(cmp)
}
