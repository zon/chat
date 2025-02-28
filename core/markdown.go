package core

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

var sanitizerPolicy *bluemonday.Policy

func GetSanitizer() *bluemonday.Policy {
	if sanitizerPolicy == nil {
		sanitizerPolicy = bluemonday.UGCPolicy()
	}
	return sanitizerPolicy
}

func MarkdownToHtml(md string) string {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	flags := html.CommonFlags
	opts := html.RendererOptions{Flags: flags}
	renderer := html.NewRenderer(opts)
	result := markdown.Render(doc, renderer)

	policy := GetSanitizer()
	return string(policy.SanitizeBytes(result))
}
