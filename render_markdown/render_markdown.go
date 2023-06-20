package render_markdown

import (
	"bytes"
	"fmt"
	"os"

	mermaid "github.com/abhinav/goldmark-mermaid"
	// "github.com/abhinav/goldmark-toc"
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"

	admonitions "github.com/stefanfritsch/goldmark-admonitions"
	fences "github.com/stefanfritsch/goldmark-fences"
	toc "github.com/stefanfritsch/goldmark-toc"
)

func RenderMarkdown(markdown []byte) string {
	// replace Windows-style line endings
	// cleanedContent := strings.ReplaceAll(string(textContent[:]), "\r\n", "\n")

	// render content
	var htmlContent bytes.Buffer

	parseMarkdown(markdown, &htmlContent)

	return htmlContent.String()
}

func RenderMarkdownFile(path string) string {
	textContent, err := os.ReadFile(path)
	
	textContent = preprocessContent(textContent)

	if err != nil {
		fmt.Printf("error loading file: %v", err)
	}

	return RenderMarkdown(textContent)
}

func parseMarkdown(source []byte, output *bytes.Buffer) {
	gm := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			extension.GFM,
			extension.DefinitionList,
			extension.Typographer,
			&fences.Extender{},
			&admonitions.Extender{},
			&toc.Extender{AddFences: true, PruneTOC: true},
			&mermaid.Extender{},
			mathjax.MathJax,
			highlighting.Highlighting,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithAttribute(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	reader := text.NewReader(source)
	pc := parser.NewContext()
	doc := gm.Parser().Parse(reader, parser.WithContext(pc))

	if err := gm.Renderer().Render(output, source, doc); err != nil {
		fmt.Printf("markdown conversion error: %v", err)
	}
}
