package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdown(t *testing.T) {
	md := "a\nb\nc"
	html := "<p>a<br>\nb<br>\nc</p>\n"
	res, err := MarkdownToHtml(md)
	assert.NoError(t, err)
	assert.Equal(t, html, res)
}

func TestMarkdownCode(t *testing.T) {
	md := "```\na\nb\nc\n```"
	html := "<pre><code>a\nb\nc\n</code></pre>\n"
	res, err := MarkdownToHtml(md)
	assert.NoError(t, err)
	assert.Equal(t, html, res)
}
