package main

import (
	"net/http"
	"strings"

	"eagle_site_v2/page"
	"eagle_site_v2/tmpl"
)

// fileServerFilter blocks directory listings under static-asset paths.
func fileServerFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") && r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func addRoutes(mux *http.ServeMux, r *tmpl.Renderer) {
	// Static content (CSS, JS, images) served from ./content.
	assets := http.FileServer(http.Dir("content"))
	mux.Handle("GET /content/", http.StripPrefix("/content", fileServerFilter(assets)))

	// Pages.
	mux.Handle("GET /", page.HomeHandler(r))
	mux.Handle("GET /hangar", page.HangarHandler(r))
	mux.Handle("GET /scenarios", page.ScenariosHandler(r))
	mux.Handle("GET /cooperation", page.CooperationHandler(r))
	mux.Handle("GET /crew", page.CrewHandler(r))
	mux.Handle("GET /contact", page.ContactGetHandler(r))
	mux.Handle("POST /contact", page.ContactPostHandler(r))
}
