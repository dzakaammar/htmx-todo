package static

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed dist
var dist embed.FS

func Handler(mux *chi.Mux) {
	mux.Handle("/*", http.StripPrefix("/",
		http.FileServer(http.FS(dist)),
	))
}
