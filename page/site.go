package page

import (
	"log/slog"
	"net/http"

	"eagle_site_v2/tmpl"
)

func renderOrFail(w http.ResponseWriter, r *tmpl.Renderer, key string, data any) {
	if err := r.Render(w, key, data); err != nil {
		slog.Error("render failed", slog.String("page", key), slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

func newBaseData(lang tmpl.Lang, key, bodyClass, navKey, reqPath string,
	content any, flash *tmpl.FlashMessage) tmpl.BasePageData {
	site := tmpl.SiteFor(lang)
	meta := site.Meta[key]
	return tmpl.BasePageData{
		Lang:         lang,
		Title:        meta.Title,
		MetaDesc:     meta.MetaDesc,
		ActiveNavKey: navKey,
		BodyClass:    bodyClass,
		HomeHref:     tmpl.HomeHref(lang),
		NavItems:     site.NavItems,
		Strings:      site.Strings,
		LangLinks:    tmpl.BuildLangLinks(reqPath, lang),
		Flash:        flash,
		Content:      content,
	}
}

func HomeHandler(r *tmpl.Renderer, lang tmpl.Lang) http.HandlerFunc {
	site := tmpl.SiteFor(lang)
	return func(w http.ResponseWriter, req *http.Request) {
		renderOrFail(w, r, "home",
			newBaseData(lang, "home", "page-home", "home", req.URL.Path, site.Home, nil))
	}
}

func HangarHandler(r *tmpl.Renderer, lang tmpl.Lang) http.HandlerFunc {
	site := tmpl.SiteFor(lang)
	return func(w http.ResponseWriter, req *http.Request) {
		renderOrFail(w, r, "hangar",
			newBaseData(lang, "hangar", "page-hangar", "hangar", req.URL.Path, site.Hangar, nil))
	}
}

func ScenariosHandler(r *tmpl.Renderer, lang tmpl.Lang) http.HandlerFunc {
	site := tmpl.SiteFor(lang)
	return func(w http.ResponseWriter, req *http.Request) {
		renderOrFail(w, r, "scenarios",
			newBaseData(lang, "scenarios", "page-scenarios", "scenarios", req.URL.Path, site.Scenarios, nil))
	}
}

func CooperationHandler(r *tmpl.Renderer, lang tmpl.Lang) http.HandlerFunc {
	site := tmpl.SiteFor(lang)
	return func(w http.ResponseWriter, req *http.Request) {
		renderOrFail(w, r, "cooperation",
			newBaseData(lang, "cooperation", "page-cooperation", "cooperation", req.URL.Path, site.Cooperation, nil))
	}
}

func CrewHandler(r *tmpl.Renderer, lang tmpl.Lang) http.HandlerFunc {
	site := tmpl.SiteFor(lang)
	return func(w http.ResponseWriter, req *http.Request) {
		renderOrFail(w, r, "crew",
			newBaseData(lang, "crew", "page-crew", "crew", req.URL.Path, site.Crew, nil))
	}
}
