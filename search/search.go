package search

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/blevesearch/bleve/v2"
)

var reBuchstaben = regexp.MustCompile(`[^\W_]$`)

func Search(index bleve.Index, queryString string) ([]byte, error) {
	if len(queryString) == 0 {
		return nil, fmt.Errorf("empty query string")
	}
	if reBuchstaben.MatchString(queryString) {
		queryString = queryString + "*"
	}
	// search for some text
	query := bleve.NewQueryStringQuery(queryString)
	search := bleve.NewSearchRequest(query)
	search.Highlight = bleve.NewHighlightWithStyle("html")
	search.Fields = []string{"DisplayName"}
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(searchResults.Hits) == 0 {
		fuzzyQuery := bleve.NewFuzzyQuery(queryString)
		search = bleve.NewSearchRequest(fuzzyQuery)
		search.Highlight = bleve.NewHighlightWithStyle("html")
		search.Fields = []string{"DisplayName"}
		searchResults, err = index.Search(search)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	if len(searchResults.Hits) > 0 {
		var res []map[string]interface{}

		for i, hit := range searchResults.Hits {

			res =
				append(res,
					map[string]interface{}{
						"order":  i,
						"path":   strings.TrimLeft(hit.ID, "docs"),
						"name":   hit.Fields["DisplayName"].(string),
						"sample": hit.Fragments["Content"],
					})
		}

		resJson, err := json.Marshal(res)

		if err != nil {
			return nil, err
		}

		return resJson, nil
	}

	return nil, fmt.Errorf("no results")
}

type searchRequest struct {
	Term string `json:"term"`
}
