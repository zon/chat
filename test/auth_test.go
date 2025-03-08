package server_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zon/hxcore"
	"github.com/zon/hxcore/test"
	idtest "github.com/zon/id/test"
)

func TestAuth(t *testing.T) {
	c := idtest.Auth(t)

	res, _ := test.HxPost(t, c.Action, map[string]string{"token": c.Token})
	test.AssertOk(t, res)
	assert.True(t, strings.HasPrefix(res.Request.URL.Path, "/user/"), "Expected /user/ prefix, got %s", res.Request.URL.Path)

	name := hxcore.RandomString(16)
	res, doc := test.HxPost(t, res.Request.URL.String(), map[string]string{"name": name})
	test.AssertOk(t, res)
	test.AssertNoErrorMsg(t, doc)
	assert.Equal(t, "/", test.HxPushUrl(res))
	assert.Equal(t, name, doc.Find("#user").Text())
}
