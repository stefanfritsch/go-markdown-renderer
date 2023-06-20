package parse_dir

import (
	"fmt"
	"strings"
)

const TocRoot = `
		<h1 class="tree-toc-root title level%v %s">%s</h1>
		<ul class="tree-toc-subdir children" style="display: %s;">
			%s
			%s
		</ul>
`

const TocEntry = `
	<li class="tree-toc-subdir subdir-item level%v">
		<div class="tree-toc-subdir title subdir-button %s">%s</button></div>
		<ul class="tree-toc-subdir children subdir-panel" style="display: %s;">
			%s
			%s
		</ul>
	</li>
`

func RenderTreeToc(dir IDir, templateSelf string, templateChildren string, pathArray []string) string {
	// will be changed to link to a file with the same name as the dir
	selfLink := fmt.Sprintf(`<div class="subdir-name">%s</div>`, dir.DisplayName())
	selfLinkSet := false

	lvl := dir.Level()
	var visibility string
	if lvl <= 0 || (lvl < len(pathArray) && pathArray[lvl] == dir.BaseName()) {
		visibility = "block"
	} else {
		visibility = "none"
	}

	filesOutput := ""

	for _, key := range dir.FileOrder() {
		file := dir.Files()[key]

		addClass := ""
		if lvl+2 == len(pathArray) && pathArray[lvl+1] == file.BaseName() {
			addClass = "active"
		}

		path := "/" + strings.TrimPrefix(file.Path(), "docs/")

		if (dir.DisplayName() == file.DisplayName() || file.DisplayName() == "README") && !selfLinkSet {
			selfLink = fmt.Sprintf(`<a class="nav-link %s" href="%s" data-path="%s">%s</a>`, addClass, path, path, dir.DisplayName())
			selfLinkSet = true
		} else {
			filesOutput = filesOutput + "\n" + fmt.Sprintf(`<li><a class="nav-link %s" href="%s" data-path="%s">%s</a></li>`, addClass, path, path, file.DisplayName())
		}
	}

	dirsOutput := ""
	if len(dir.Subdirs()) > 0 {
		for _, key := range dir.SubdirOrder() {
			if dir.BaseName() == "assets" || dir.BaseName() == "deps" {
				return ""
			}
			subdir := dir.Subdirs()[key]
			dirsOutput = dirsOutput + RenderTreeToc(subdir, templateChildren, templateChildren, pathArray)
		}
	}

	addClass := ""
	if visibility == "block" {
		addClass = "active"
	}

	output := fmt.Sprintf(templateSelf, lvl, addClass, selfLink, visibility, filesOutput, dirsOutput)

	return output
}
