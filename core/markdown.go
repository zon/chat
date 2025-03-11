package core

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

func GetMdConverter() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)
}

func MarkdownToHtml(md string) (string, error) {
	text := strings.ReplaceAll(md, "<br>", "\n")
	var result bytes.Buffer
	err := GetMdConverter().Convert([]byte(text), &result)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
