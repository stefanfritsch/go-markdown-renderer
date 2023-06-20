package backlinks

import (
	"go-markdown-renderer/parse_dir"
	"go-markdown-renderer/render_markdown"
	"path"
)

func ParseBacklinks(parsedDir parse_dir.IDir, markdownExts []string) *map[string][]parse_dir.IEntry {
	backlinks := map[string][]parse_dir.IEntry{}

	parse_dir.WalkIf(
		parsedDir,
		func(e parse_dir.IEntry) bool {
			ext := parse_dir.FileExt(e.BaseName())
			res := false

			for _, e := range markdownExts {
				if ext == e {
					res = true
				}
			}
			return res
		},
		func(e parse_dir.IEntry) {
			ctx := render_markdown.ParseLinksAndYamlFile(e.Path())

			for _, ref := range ctx.References() {
				lbl := string(ref.Label())
				if lbl == "//begin" || lbl == "//end" {
					continue
				}

				fullRelPath := path.Join(e.DirName(), string(ref.Destination()))

				currEntries, exists := backlinks[fullRelPath]
				if !exists {
					backlinks[fullRelPath] = []parse_dir.IEntry{e}
				} else {
					backlinks[fullRelPath] = append(currEntries, e)
				}
			}
		})

	return &backlinks
}
