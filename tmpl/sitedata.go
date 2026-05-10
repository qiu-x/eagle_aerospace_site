package tmpl

import "strings"

type SiteData struct {
	NavItems    []NavItem
	Strings     map[string]string
	Meta        map[string]PageMeta
	Home        HomePageData
	Hangar      HangarPageData
	Scenarios   ScenariosPageData
	Cooperation CooperationPageData
	Crew        CrewPageData
	Contact     ContactPageData
}

// Returns the SiteData for the given language. PL is the fallback.
func SiteFor(l Lang) SiteData {
	if l == LangEN {
		return siteEN
	}
	return sitePL
}

// Localized site root, e.g. "/pl/" or "/en/".
func HomeHref(l Lang) string {
	return "/" + string(l) + "/"
}

func SwitchPath(path string, to Lang) string {
	for _, p := range []string{"/" + string(LangPL), "/" + string(LangEN)} {
		if path == p || strings.HasPrefix(path, p+"/") {
			return "/" + string(to) + path[len(p):]
		}
	}
	return HomeHref(to)
}

func BuildLangLinks(path string, current Lang) []LangLink {
	return []LangLink{
		{Code: LangPL, Label: "PL", Href: SwitchPath(path, LangPL), Active: current == LangPL},
		{Code: LangEN, Label: "EN", Href: SwitchPath(path, LangEN), Active: current == LangEN},
	}
}
