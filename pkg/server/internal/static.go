package internal

import (
	"embed"
	"io/fs"
)

//go:embed static
var staticFiles embed.FS

func GetStaticFiles() fs.FS {
	subFS, _ := fs.Sub(staticFiles, "static")
	return subFS
}
