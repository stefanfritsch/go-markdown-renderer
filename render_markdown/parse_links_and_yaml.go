package render_markdown

import (
	"fmt"
	"os"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)


func ParseLinksAndYaml(source []byte) parser.Context {
	p := parser.NewParser(
		parser.WithBlockParsers(
			util.Prioritized(parser.NewParagraphParser(), 1000),
		),
		parser.WithInlineParsers(
			util.Prioritized(parser.NewLinkParser(), 200),
			util.Prioritized(parser.NewAutoLinkParser(), 300),
		),
		parser.WithParagraphTransformers(
			util.Prioritized(parser.LinkReferenceParagraphTransformer, 100),
		),
	)

	gm := goldmark.New(
		goldmark.WithParser(p),
		goldmark.WithExtensions(
			meta.Meta,
		),
	)
	reader := text.NewReader(source)

	parseContext := parser.NewContext()
	gm.Parser().Parse(reader, parser.WithContext(parseContext))

	return parseContext
}

func ParseLinksAndYamlFile(path string) parser.Context {
	textContent, err := os.ReadFile(path)

	if err != nil {
		fmt.Printf("error loading file: %v", err)
	}

	return ParseLinksAndYaml(textContent)
}
