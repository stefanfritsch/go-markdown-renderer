package parse_dir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBaseName(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		"path",
		GetBaseName("/this/is/a/path"),
		"Basic functionality",
	)

	assert.Equal(
		"path",
		GetBaseName("/this/is/a/path/"),
		"Trailing slash",
	)

	assert.Equal(
		"path",
		GetBaseName("this/is/a/path/"),
		"No leading slash",
	)

	assert.Equal(
		"path",
		GetBaseName("/this/is/a/path//"),
		"Multiple trailing slashes",
	)

	assert.Equal(
		"",
		GetBaseName("/"),
		"Only a slash",
	)

	assert.Equal(
		"path",
		GetBaseName("path"),
		"No slashes",
	)

}

func TestGetDirName(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		"/this/is/a",
		GetDirName("/this/is/a/path"),
		"Basic functionality",
	)

	assert.Equal(
		"/this/is/a",
		GetDirName("/this/is/a/path/"),
		"Trailing slash",
	)

	assert.Equal(
		"this/is/a",
		GetDirName("this/is/a/path/"),
		"No leading slash",
	)

	assert.Equal(
		"/this/is/a",
		GetDirName("/this/is/a/path//"),
		"Multiple trailing slashes",
	)

	assert.Equal(
		"",
		GetDirName("/"),
		"Only a slash",
	)

	assert.Equal(
		"",
		GetDirName("path"),
		"No slashes",
	)

}
