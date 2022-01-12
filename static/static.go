package static

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed content
var content embed.FS

func RegisterStaticContent(r *mux.Router, prefix string) {
	fsys, err := fs.Sub(content, "content")
	if err != nil {
		log.Printf("failed to register static content: %v", err)

		return
	}

	r.NewRoute().PathPrefix(prefix).Handler(
		http.StripPrefix(prefix,
			http.FileServer(http.FS(fsys))))
}
