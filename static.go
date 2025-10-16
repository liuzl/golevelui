package golevelui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:static
var embeddedFS embed.FS

// getStaticFS returns a filesystem (http.FileSystem) for the embedded static assets.
// It creates a sub-filesystem rooted at the "static" directory, which allows the
// http.FileServer to serve files correctly.
func getStaticFS() http.FileSystem {
	subFS, err := fs.Sub(embeddedFS, "static")
	if err != nil {
		// This should not happen with a correctly embedded directory
		panic(err)
	}
	return http.FS(subFS)
}
