// Package page holds the HTTP handlers that render every page of the
// Eagle Aerospace marketing site. Each handler is a thin adapter that
// assembles a tmpl.BasePageData and delegates to the renderer.
package page

import (
	"log/slog"
	"net/http"

	"eagle_site_v2/tmpl"
)

// renderOrFail is the common tail of every GET page handler. It logs
// and returns a generic 500 if the renderer fails (which only happens
// for programmer errors at this point, since templates were parsed at
// boot).
func renderOrFail(w http.ResponseWriter, r *tmpl.Renderer, key string, data any) {
	if err := r.Render(w, key, data); err != nil {
		slog.Error("render failed", slog.String("page", key), slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

func HomeHandler(r *tmpl.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Catch any path under "/" that wasn't routed elsewhere.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		renderOrFail(w, r, "home", tmpl.BasePageData{
			Title:        "Eagle Aerospace - Systemy Bezzałogowe",
			MetaDesc:     "Polska firma projektująca bezzałogowe statki powietrzne dla sektora publicznego, służb mundurowych i instytucji badawczych.",
			ActiveNavKey: "home",
			BodyClass:    "page-home",
			NavItems:     tmpl.NavItems,
			Content:      tmpl.HomeContent,
		})
	}
}

func HangarHandler(r *tmpl.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		renderOrFail(w, r, "hangar", tmpl.BasePageData{
			Title:        "Hangar - Eagle Aerospace",
			MetaDesc:     "Katalog platform bezzałogowych Eagle Aerospace - płatowce, VTOL, wielowirnikowce.",
			ActiveNavKey: "hangar",
			BodyClass:    "page-hangar",
			NavItems:     tmpl.NavItems,
			Content:      tmpl.HangarContent,
		})
	}
}

func ScenariosHandler(r *tmpl.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		renderOrFail(w, r, "scenarios", tmpl.BasePageData{
			Title:        "Scenariusze - Eagle Aerospace",
			MetaDesc:     "Inspekcje techniczne, rolnictwo precyzyjne, ratownictwo, monitoring środowiska - operacyjne zastosowania platform Eagle Aerospace.",
			ActiveNavKey: "scenarios",
			BodyClass:    "page-scenarios",
			NavItems:     tmpl.NavItems,
			Content:      tmpl.ScenariosContent,
		})
	}
}

func CooperationHandler(r *tmpl.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		renderOrFail(w, r, "cooperation", tmpl.BasePageData{
			Title:        "Współpraca - Eagle Aerospace",
			MetaDesc:     "Współpraca Eagle Aerospace z uczelniami, jednostkami badawczymi i instytucjami publicznymi.",
			ActiveNavKey: "cooperation",
			BodyClass:    "page-cooperation",
			NavItems:     tmpl.NavItems,
			Content:      tmpl.CooperationContent,
		})
	}
}

func CrewHandler(r *tmpl.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		renderOrFail(w, r, "crew", tmpl.BasePageData{
			Title:        "Za sterami - Eagle Aerospace",
			MetaDesc:     "Zespół inżynierów i konstruktorów Eagle Aerospace.",
			ActiveNavKey: "crew",
			BodyClass:    "page-crew",
			NavItems:     tmpl.NavItems,
			Content:      tmpl.CrewContent,
		})
	}
}
