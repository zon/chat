package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func markdownToHtml(md string) string {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	flags := html.CommonFlags
	opts := html.RendererOptions{Flags: flags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}
