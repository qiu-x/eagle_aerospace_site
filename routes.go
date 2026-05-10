package main

import (
	"net/http"
	"strings"

	"eagle_site_v2/page"
	"eagle_site_v2/tmpl"
)

// blocks directory listings under static-asset paths.
func fileServerFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") && r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func addLocalizedRoutes(mux *http.ServeMux, prefix string, lang tmpl.Lang, r *tmpl.Renderer) {
	mux.Handle("GET "+prefix+"/{$}", page.HomeHandler(r, lang))
	mux.Handle("GET "+prefix+"/hangar", page.HangarHandler(r, lang))
	mux.Handle("GET "+prefix+"/scenarios", page.ScenariosHandler(r, lang))
	mux.Handle("GET "+prefix+"/cooperation", page.CooperationHandler(r, lang))
	mux.Handle("GET "+prefix+"/crew", page.CrewHandler(r, lang))
	mux.Handle("GET "+prefix+"/contact", page.ContactGetHandler(r, lang))
	mux.Handle("POST "+prefix+"/contact", page.ContactPostHandler(r, lang))
}

func addRoutes(mux *http.ServeMux, r *tmpl.Renderer) {
	assets := http.FileServer(http.Dir("content"))
	mux.Handle("GET /content/", http.StripPrefix("/content", fileServerFilter(assets)))

	addLocalizedRoutes(mux, "/pl", tmpl.LangPL, r)
	addLocalizedRoutes(mux, "/en", tmpl.LangEN, r)

	// 302 REDIRECT to a localized home based on Accept-Language.
	mux.HandleFunc("GET /{$}", page.LocaleRedirect)

	for _, slug := range []string{"hangar", "scenarios", "cooperation", "crew", "contact"} {
		target := "/pl/" + slug
		mux.HandleFunc("GET /"+slug, func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, target, http.StatusMovedPermanently)
		})
	}
}
