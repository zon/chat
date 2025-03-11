package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdown(t *testing.T) {
	md := `a
b
c`
	html := `<p>a<br>
b<br>
c</p>
`
	res := MarkdownToHtml(md)
	assert.Equal(t, html, res)
}

func TestMarkdownCode(t *testing.T) {
	md := "```a\nb\nc```"
	html := `<code>a
b
b</code>`
	res := MarkdownToHtml(md)
	assert.Equal(t, html, res)
}