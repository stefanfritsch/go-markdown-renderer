package render_markdown

import "regexp"

var detectFoamLinkWithLabel = regexp.MustCompile(`\[\[([^|\]]*?)\|([^|\]]*?)\]\]`)

func preprocessContent(textContent []byte) []byte {
	output := detectFoamLinkWithLabel.ReplaceAll(textContent, []byte("[[$2][$1|$2]]"))
	
	return output
}
