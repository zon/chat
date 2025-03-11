package core

import (
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

var textSanitizer *bluemonday.Policy
var htmlSanitizer *bluemonday.Policy
var markdownRenderer *html.Renderer

func GetTextSanitizer() *bluemonday.Policy {
	if textSanitizer == nil {
		textSanitizer = bluemonday.StrictPolicy()
	}
	return textSanitizer
}

func GetHtmlSanitizer() *bluemonday.Policy {
	if htmlSanitizer == nil {
		htmlSanitizer = bluemonday.UGCPolicy()
	}
	return htmlSanitizer
}

func GetMarkdownParser() *parser.Parser {
	extensions := parser.CommonExtensions | parser.HardLineBreak
	return parser.NewWithExtensions(extensions)
}

func GetMarkdownRenderer() *html.Renderer {
	if markdownRenderer == nil {
		flags := html.CommonFlags
		opts := html.RendererOptions{Flags: flags}
		markdownRenderer = html.NewRenderer(opts)
	}
	return markdownRenderer
}

func MarkdownToHtml(md string) string {
	text := strings.ReplaceAll(md, "<br>", "\n")
	text = string(GetTextSanitizer().Sanitize(text))

	fmt.Println("text: ", text)

	result := markdown.ToHTML(
		[]byte(md),
		GetMarkdownParser(),
		GetMarkdownRenderer(),
	)

	fmt.Println("result: ", string(result))

	return string(GetHtmlSanitizer().SanitizeBytes(result))
}
