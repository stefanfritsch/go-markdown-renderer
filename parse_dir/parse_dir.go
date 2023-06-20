package parse_dir

import (
	"go-markdown-renderer/helpers"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strings"
)

func ParseDir(dirpath string, filePattern string) IDir {
	currDir := strings.TrimSuffix(dirpath, "/")
	pattern := regexp.MustCompile(filePattern)

	entry := &Dir{
		Entry: Entry{
			displayName: currDir,
			path:        currDir,
			parent:      &Dir{},
			level:       0,
		},
		subdirs: map[string]IDir{},
		files:   map[string]IFile{},
	}

	dir := parseRecursively(entry, pattern)

	return dir
}

func parseRecursively(dir IDir, pattern *regexp.Regexp) IDir {
	parsedDir := parseSubDir(dir, pattern)

	for _, subdir := range parsedDir.Subdirs() {
		parseRecursively(subdir, pattern)
	}

	return parsedDir
}

func parseSubDir(dir IDir, pattern *regexp.Regexp) IDir {
	entries, err := os.ReadDir(dir.Path())
	if err != nil {
		log.Fatal(err)
	}

	subdirs := make(map[string]IDir)
	subdirOrder := []string{}
	files := make(map[string]IFile, len(entries))
	fileOrder := []string{}

	for _, entry := range entries {
		// skip hidden files
		if entry.Name()[0] == []byte(".")[0] ||
		   entry.Name()[0] == []byte("_")[0] {
			continue
		}

		// create dir and file entries
		if entry.IsDir() {
			subdirs[entry.Name()] = createDir(entry, dir)
			subdirOrder = append(subdirOrder, entry.Name())
		} else {
			if pattern.MatchString(entry.Name()) {
				files[entry.Name()] = createFile(entry, dir)
				fileOrder = append(fileOrder, entry.Name())
			}
		}
	}

	helpers.SortAlphabetically(subdirOrder)
	helpers.SortAlphabetically(fileOrder)

	dir.Subdirs(subdirs)
	dir.SubdirOrder(subdirOrder)
	dir.Files(files)
	dir.FileOrder(fileOrder)

	return dir
}

func createDir(entry fs.DirEntry, parent IDir) IDir {
	return &Dir{
		Entry: Entry{
			displayName: entry.Name(),
			path:        parent.Path() + "/" + entry.Name(),
			parent:      parent,
			level:       parent.Level() + 1,
		},
		subdirs: map[string]IDir{},
		files:   map[string]IFile{},
	}
}

func createFile(entry fs.DirEntry, parent IDir) IFile {
	return &File{
		Entry{
			displayName: FileWithoutExt(entry.Name()),
			path:        parent.Path() + "/" + entry.Name(),
			parent:      parent,
			level:       parent.Level() + 1,
		},
	}
}
