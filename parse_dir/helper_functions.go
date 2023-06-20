package parse_dir

// find the position of the last slash in the path or -1 if there is none.
func posLastSlash(path string) int {
	start := -1

	for i := 0; i < len(path)-1; i++ {
		if path[i] == []byte("/")[0] {
			start = i
		}
	}

	return start
}

// Return the path from the last slash
//
// ### Description
//
// This function returns the path from the last slash after trailing slashes
// have been removed. E.g. "/this/is/a/path/" returns "path".
//
// If there is no slash, the full input path is returned.
func GetBaseName(path string) string {
	// if path ends with "/" and is not just a "/", remove that trailing slash
	for path[len(path)-1] == []byte("/")[0] {
		if len(path) <= 1 {
			return ""
		}

		path = path[:len(path)-1]
	}

	startLastName := posLastSlash(path)

	if startLastName < 0 {
		return path
	}

	return path[startLastName+1:]
}

// Return the path up to the last slash
//
// ### Description
//
// This function returns a path up to the last slash after trailing slashes
// have been removed. e.g. "/this/is/a/path/" would return "/this/is/a".
//
// If there is no slash, an empty string is returned. If the path is the root
// (i.e. "/") an empty string is returned as well.
func GetDirName(path string) string {
	// if path ends with "/" and is not just a "/", remove that trailing slash
	if path[len(path)-1] == []byte("/")[0] {
		if len(path) == 1 {
			return ""
		}

		path = path[:len(path)-1]
	}

	startLastName := posLastSlash(path)

	if startLastName < 1 {
		return ""
	}

	return path[:startLastName]
}

// find the start of the file extension
// this is the last "." in the filename if the filename is at least 3 characters
// long (which should exclude . and .. but not a.b)
func beginFileExt(s string) int {
	if len(s) < 3 {
		return -1
	}

	start := -1

	for i := 1; i < len(s)-1; i++ {
		if s[i] == []byte(".")[0] {
			start = i
		}
	}

	return start
}

// get the file extension (i.e. elements from the last "." or "" if there is no "." in the name)
func FileExt(s string) string {
	startExt := beginFileExt(s)

	if startExt < 1 {
		return ""
	}

	return s[startExt:]
}

// get the file name without the file extension
func FileWithoutExt(s string) string {
	startExt := beginFileExt(s)

	if startExt < 1 {
		return s
	}

	return s[:startExt]
}

// Apply the function f to all entries of d and its subdirs
func Walk(d IDir, f func(IEntry)) {
	p := func(IEntry) bool { return true }

	WalkIf(d, p, f)
}

// Apply the function f to all entries of d and its subdirs if p returns true
func WalkIf(d IDir, p func(IEntry) bool, f func(IEntry)) {
	for _, file := range d.Files() {
		if p(file) {
			f(file)
		}
	}

	for _, subdir := range d.Subdirs() {
		WalkIf(subdir, p, f)
	}
}
