package admin

import (
	"net/http"
	"os"
	"path/filepath"
)

// spaFileSystem is a custom file system that serves the index.html file
// for any path that does not have a file extension and does not correspond
// to an existing file.
type spaFileSystem struct {
	root      http.FileSystem
	indexPath string
}

// Open opens the file at the given path. If the file is not found, and the path
// does not have a file extension, it serves the configured index file.
func (fs *spaFileSystem) Open(name string) (http.File, error) {
	f, err := fs.root.Open(name)
	// If the file doesn't exist, and it's not a request for an asset (judged by file extension),
	// serve the index.html page.
	if os.IsNotExist(err) && filepath.Ext(name) == "" {
		return fs.root.Open(fs.indexPath)
	}
	return f, err
}

// SpaHandler returns an http.Handler that serves a Single Page Application.
func SpaHandler(staticPath, indexPath string) http.Handler {
	return http.FileServer(&spaFileSystem{
		root:      http.Dir(staticPath),
		indexPath: indexPath,
	})
}
