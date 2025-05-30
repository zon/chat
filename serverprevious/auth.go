package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zon/chat/core"
)

const ZitadelClientID = "322307909566444934"

const AuthTokenSecretEnv = "AUTH_TOKEN_SECRET"
const DefaultAuthTokenSecret = "beans"

var AuthTokenSecret []byte

type postAuthBody struct {
	Token string
}

type authTokenClaims struct {
	jwt.RegisteredClaims
	UserID uint
}

func postAuth(c *fiber.Ctx) error {
	var body postAuthBody
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	var claims authTokenClaims
	_, err = jwt.ParseWithClaims(body.Token, &claims, authTokenKey)
	if err != nil {
		return err
	}

	if claims.UserID == 0 {
		return fmt.Errorf("no userID in token")
	}

	session, err := core.GetSession(c)
	if err != nil {
		return err
	}
	err = session.SetUserID(c, claims.UserID)
	if err != nil {
		return err
	}

	return c.Redirect("/")
}

func authTokenKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return AuthTokenSecret, nil
}

func LoadAuthTokenSecret() {
	AuthTokenSecret = []byte(os.Getenv(AuthTokenSecretEnv))
	if len(AuthTokenSecret) < 1 {
		AuthTokenSecret = []byte(DefaultAuthTokenSecret)
	}
}
