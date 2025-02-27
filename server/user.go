package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zon/chat/core"
	"github.com/zon/chat/html"
)

type postUserBody struct {
	Name string
}

func getUser(c *fiber.Ctx) error {
	session, err := core.GetSession(c)
	if err != nil {
		return err
	}
	user, err := core.GetUser(session.UserID)
	if err != nil {
		return err
	}

	layout := html.Layout(html.User(user))
	if isHxRequest(c) {
		return render(c, layout)
	} else {
		return render(c, html.Doc(layout))
	}
}

func postUser(c *fiber.Ctx) error {
	session, err := core.GetSession(c)
	if err != nil {
		return err
	}
	user, err := core.GetUser(session.UserID)
	if err != nil {
		return err
	}

	var body postUserBody
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}
	name := body.Name

	if name == "" {
		return renderError(c, "Name required")
	}

	user.Name = name
	err = user.Save()
	if err != nil {
		return err
	}

	session.Ready = true
	err = session.Save(c)
	if err != nil {
		return err
	}

	return redirectToIndex(c, "Name changed")
}
