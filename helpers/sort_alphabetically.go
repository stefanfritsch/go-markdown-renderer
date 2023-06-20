package helpers

import (
	"sort"
	"strings"
	"unicode"
)

func SortAlphabetically(s []string) {
	sort.Slice(s, func(i, j int) bool {
		t := strings.Map(unicode.ToUpper, s[i])
		u := strings.Map(unicode.ToUpper, s[j])
		return t < u
	})
}
