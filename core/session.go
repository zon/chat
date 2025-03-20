package core

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memcache"
	"github.com/gofiber/storage/sqlite3"
)

const userKey string = "userID"
const readyKey string = "ready"

var store *session.Store

type Session struct {
	UserID uint
	Ready  bool
	Fiber  *session.Session
}

func (s *Session) ID() string {
	return s.Fiber.ID()
}

func (s *Session) HasAuth() bool {
	return s.UserID > 0
}

func (s *Session) SetUserID(ctx *fiber.Ctx, id uint) error {
	user, err := GetUser(id)
	if err != nil {
		return err
	}
	s.UserID = user.ID
	s.Ready = user.Ready
	return s.Save(ctx)
}

func (s *Session) Save(ctx *fiber.Ctx) error {
	s.Fiber.Set(userKey, s.UserID)
	s.Fiber.Set(readyKey, s.Ready)
	return s.Fiber.Save()
}

func InitSessionStore() {
	host := os.Getenv("MEMCACHED_HOST")
	port := os.Getenv("MEMCACHED_PORT")

	var storage fiber.Storage
	if host != "" {
		storage = memcache.New(memcache.Config{
			Servers: fmt.Sprintf("%s:%s", host, port),
		})

	} else {
		storage = sqlite3.New()
	}

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
	session.Fiber = fs
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
