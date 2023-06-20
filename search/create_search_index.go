package search

import (
	"fmt"
	"go-markdown-renderer/parse_dir"
	"os"
	"regexp"

	"github.com/blevesearch/bleve/v2"
)

func CreateSearchIndex(indexPath string, parsedDir parse_dir.IDir) (bleve.Index, error) {
	// See here: https://github.com/blevesearch/bleve/issues/1711
	// for the source of the code
	mapping := bleve.NewIndexMapping()
	mapping.DefaultAnalyzer = "standard"

	// if err := mapping.AddCustomAnalyzer("custom-html", map[string]interface{}{
	// 	"type":         bleveCustom.Name,
	// 	"tokenizer":    bleveWeb.Name,
	// 	"char_filters": []interface{}{bleveHtml.Name},
	// }); err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }

	// fm := bleveMapping.NewTextFieldMapping()
	// fm.Analyzer = "custom-html"

	// dmap := bleveMapping.NewDocumentMapping()
	// dmap.AddFieldMappingsAt("Content", fm)

	// mapping.DefaultMapping = dmap

	index, err := bleve.New(indexPath, mapping)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	r := regexp.MustCompile(`(?s)\[//begin\].*?\[//end\]`)

	f := func(e parse_dir.IEntry) {
		content, err := os.ReadFile(e.Path())
		content = r.ReplaceAll(content, []byte(""))
		// content := render_markdown.RenderMarkdownFile(e.Path())
		// content = `<div id="mw-page-base" class="noprint"> found </div>`
		if err != nil {
			panic(err)
		}

		index.Index(e.Path(), map[string]string{
			"DisplayName": e.DisplayName(),
			"Content":     string(content),
		})
	}

	// index some data
	parse_dir.Walk(parsedDir, f)

	return index, nil
}
