package assets

import (
	"io/fs"
	"net/http"
)

var (
	content    fs.FS
	FileSystem http.FileSystem
	prefixPath string
)

func Load(path string) {
	prefixPath = path
	if prefixPath != "" {
		FileSystem = http.Dir(prefixPath)
	} else {
		FileSystem = http.FS(content)
	}
}

func Register(fileSystem fs.FS) {
	subFs, err := fs.Sub(fileSystem, "static")
	if err != nil {
		content = subFs
	}
}
