package page

import (
	"net/http"

	"golang.org/x/text/language"

	"eagle_site_v2/tmpl"
)

// matcher orders supported locales by preference. The first entry is
// the fallback when nothing in the Accept-Language header matches.
var matcher = language.NewMatcher([]language.Tag{
	language.Polish,
	language.English,
})

// LocaleRedirect serves GET / by negotiating Accept-Language and
// 302-redirecting to /pl/ or /en/. The URL is the only source of truth
// for locale, so this is the only place the header is consulted.
func LocaleRedirect(w http.ResponseWriter, req *http.Request) {
	target := tmpl.HomeHref(detectLang(req))
	w.Header().Set("Cache-Control", "private, no-store")
	w.Header().Set("Vary", "Accept-Language")
	http.Redirect(w, req, target, http.StatusFound)
}

// detectLang returns the best-matching supported language for the
// incoming request, defaulting to PL when the header is missing,
// empty or unparseable.
func detectLang(req *http.Request) tmpl.Lang {
	accept := req.Header.Get("Accept-Language")
	if accept == "" {
		return tmpl.LangPL
	}
	tags, _, err := language.ParseAcceptLanguage(accept)
	if err != nil || len(tags) == 0 {
		return tmpl.LangPL
	}
	matched, _, _ := matcher.Match(tags...)
	base, _ := matched.Base()
	if base.String() == "en" {
		return tmpl.LangEN
	}
	return tmpl.LangPL
}
