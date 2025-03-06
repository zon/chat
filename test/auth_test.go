package server_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zon/hxcore/test"
	idtest "github.com/zon/id/test"
)

func TestAuth(t *testing.T) {
	c := idtest.Auth(t)

	res, _ := test.HxPost(t, c.Action, map[string]string{"token": c.Token})
	test.AssertRedirect(t, "/", res)
	assert.Equal(t, "8080", res.Request.URL.Port())
	test.AssertSession(t, res)
}
