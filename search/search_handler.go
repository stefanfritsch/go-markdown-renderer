package search

import (
	"encoding/json"
	"fmt"
	"go-markdown-renderer/helpers"
	"io"
	"net/http"

	"github.com/blevesearch/bleve/v2"
	bleveHttp "github.com/blevesearch/bleve/v2/http"
)

func SearchHandleFunc(index bleve.Index) func(http.ResponseWriter, *http.Request) {

	return func(res http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			helpers.Send400(res, "No request body")
			return
		}
		var search searchRequest
		json.Unmarshal(body, &search)

		searchResults, err := Search(index, search.Term)

		if err != nil {
			helpers.Send500(res, fmt.Sprintf("%v", err))
		}

		res.Write(searchResults)
	}
}

func SearchHandler(index bleve.Index) *bleveHttp.SearchHandler {
	bleveHttp.RegisterIndexName("docs", index)
	searchHandler := bleveHttp.NewSearchHandler("docs")

	return searchHandler
}
