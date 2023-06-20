package main

import (
	"fmt"
	bl "go-markdown-renderer/backlinks"
	"go-markdown-renderer/helpers"
	"go-markdown-renderer/parse_dir"
	"go-markdown-renderer/render_markdown"
	"go-markdown-renderer/search"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// list of accepted markdown extensions. Everything else is served as application/octet-stream
var markdownExts = []string{".md", ".Md", ".MD", ".Rmd", ".rmd"}

func handleMarkdownFunc(
	basedir string,
	parsedDir parse_dir.IDir,
	markdownExts []string,
	siteTitle string,
	siteIcon string,
	backlinks *map[string][]parse_dir.IEntry,
) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		pathArray := strings.Split(path, "/")

		if path == "/" {
			path = "docs/README.md"
		} else {
			path = "docs" + path
		}

		// ========================================================================== //
		// check if the file is a markdown

		fileExt := parse_dir.FileExt(path)

		isMarkdown := false
		for _, me := range markdownExts {
			if fileExt == me {
				isMarkdown = true
			}
		}

		// ========================================================================== //
		// if it's not markdown we currently treat everything as an octet-stream

		if !isMarkdown {
			fileBytes, err := os.ReadFile(path)
			if err != nil {
				helpers.Send500(res, fmt.Sprintf("%v", err))
			}
			// res.WriteHeader(http.StatusOK)

			contentType := "application/octet-stream"
			if fileExt == ".png" {
				contentType = "image/png"
			} else if fileExt == ".html" {
				contentType = "text/html"
			} else if fileExt == ".js" {
				contentType = "text/javascript"
			} else if fileExt == ".css" {
				contentType = "text/css"
			} else if fileExt == ".json" {
				contentType = "application/json"
			} else if fileExt == ".xml" {
				contentType = "text/xml"
			} else if fileExt == ".svg" {
				contentType = "image/svg+xml"
			} else if fileExt == ".jpg" {
				contentType = "image/jpeg"
			} else if fileExt == ".gif" {
				contentType = "image/gif"
			}

			res.Header().Set("Content-Type", contentType)

			res.Write(fileBytes)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/base.html"))

		tmpl.ExecuteTemplate(res, "base", map[string]template.HTML{
			"BaseDir":   template.HTML(basedir),
			"SiteTitle": template.HTML(siteTitle),
			"SiteIcon":  template.HTML(siteIcon),
			"TreeToc":   template.HTML(parse_dir.RenderTreeToc(parsedDir, parse_dir.TocRoot, parse_dir.TocEntry, pathArray)),
			"MD":        template.HTML(render_markdown.RenderMarkdownFile(path)),
			"BackLinks": template.HTML(bl.RenderBackLinks((*backlinks)[path])),
		})
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// ========================================================================== //
	// find basedir
	var stripper (func(http.Handler) http.Handler)
	basedir := os.Getenv("BASEDIR")

	if len(basedir) == 0 {
		basedir = ""
		stripper = func(h http.Handler) http.Handler { return h }
	} else {
		if basedir[0] != '/' {
			basedir = "/" + basedir
		}
		if basedir[len(basedir)-1] == '/' {
			basedir = basedir[:len(basedir)-1]
		}
		stripper = func(h http.Handler) http.Handler { return http.StripPrefix(basedir, h) }
	}

	// ========================================================================== //
	// find port
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8090"
	}

	// ========================================================================== //
	// 	find title

	siteTitle := "Markdown Renderer"
	titleEnv := os.Getenv("TITLE")

	if len(titleEnv) > 0 {
		siteTitle = titleEnv
	}

	// ========================================================================== //
	// 	find icon
	siteIcon := os.Getenv("ICON")

	if len(siteIcon) == 0 {
		siteIcon = "fa-solid fa-book"
	}

	// ========================================================================== //
	// 	create tree

	fmt.Println("Parsing files")
	parsedDir := parse_dir.ParseDir("docs/", `.md$`)
	fmt.Println("Parsing completed")

	// ========================================================================== //
	// parse backlinks

	backlinks := bl.ParseBacklinks(parsedDir, markdownExts)

	// ========================================================================== //
	// 	create search index

	fmt.Println("Creating search index")
	idxDir := "docs.bleve"
	if _, err := os.Stat(idxDir); err == nil {
		err = os.RemoveAll(idxDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	searchIndex, err := search.CreateSearchIndex(idxDir, parsedDir)
	if err != nil {
		panic(err)
	}
	fmt.Println("Search index created")

	// ========================================================================== //
	// 	define handlers

	http.Handle(basedir+"/hello", stripper(http.HandlerFunc(hello)))
	http.Handle(basedir+"/headers", stripper(http.HandlerFunc(headers)))
	http.Handle(basedir+"/api/search", stripper(http.HandlerFunc(search.SearchHandleFunc(searchIndex))))
	http.Handle(
		basedir+"/_markdown_renderer_assets/",
		http.StripPrefix(
			basedir+"/_markdown_renderer_assets/",
			http.FileServer(http.Dir("assets/")),
		),
	)
	http.Handle(basedir+"/", stripper(http.HandlerFunc(handleMarkdownFunc(basedir, parsedDir, markdownExts, siteTitle, siteIcon, backlinks))))

	// ========================================================================== //
	// 	Listening

	fmt.Printf("Listening on http://localhost:%s%s/\n", port, basedir)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
