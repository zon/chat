package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
)

const userKey string = "userID"
const readyKey string = "ready"

var store *session.Store

type Session struct {
	UserID uint
	Ready  bool
	fiber  *session.Session
}

func CreateSession(ctx *fiber.Ctx, userID uint) (*Session, error) {
	var session *Session
	user, err := GetUser(userID)
	if err != nil {
		return session, err
	}
	session = &Session{
		UserID: user.ID,
		Ready:  user.IsReady(),
	}
	f, err := store.Get(ctx)
	if err != nil {
		return session, err
	}
	session.fiber = f
	err = session.Save(ctx)
	return session, err
}

func (s *Session) HasAuth() bool {
	return s.UserID > 0
}

func (s *Session) Save(ctx *fiber.Ctx) error {
	s.fiber.Set(userKey, s.UserID)
	s.fiber.Set(readyKey, s.Ready)
	return s.fiber.Save()
}

func InitSessionStore() {
	storage := sqlite3.New()
	store = session.New(session.Config{
		Storage: storage,
	})
}

func GetSession(ctx *fiber.Ctx) (*Session, error) {
	session := &Session{}
	fs, err := store.Get(ctx)
	if err != nil {
		return session, err
	}
	session.fiber = fs
	uv := fs.Get(userKey)
	if uv != nil {
		session.UserID = uv.(uint)
	}
	rv := fs.Get(readyKey)
	if rv != nil {
		session.Ready = rv.(bool)
	}
	return session, nil
}

func GetSessionUser(ctx *fiber.Ctx) (*User, error) {
	var user *User
	session, err := GetSession(ctx)
	if err != nil {
		return user, err
	}
	return GetUser(session.UserID)
}
