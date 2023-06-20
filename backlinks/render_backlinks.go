package backlinks

import (
	"fmt"
	"go-markdown-renderer/helpers"
	"go-markdown-renderer/parse_dir"
	"strings"
)

var backLink = `
	<a class="backlinks-backlink nav-link" href="%v">%v</a>
`

func RenderBackLinks(backlinks []parse_dir.IEntry) string {
	output := `
	<h1 id="backlinks-title">What links here?</h1>
  `

	if len(backlinks) < 1 {
		return output
	}

	links := make([]string, len(backlinks))

	for _, entry := range backlinks {
		path := "/" + strings.TrimPrefix(entry.Path(), "docs/")
		links = append(links, fmt.Sprintf(backLink, path, entry.DisplayName()))
	}

	helpers.SortAlphabetically(links)

	return output +
		`<div id="backlinks-links">` +
		strings.Join(links, "\n") +
		`</div>`
}
