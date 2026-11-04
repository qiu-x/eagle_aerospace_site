package tmpl

var sitePL = SiteData{
	NavItems: []NavItem{
		{Label: "Hangar", Href: "/pl/hangar", Key: "hangar"},
		{Label: "Scenariusze", Href: "/pl/scenarios", Key: "scenarios"},
		// {Label: "Współpraca", Href: "/pl/cooperation", Key: "cooperation"},
		{Label: "Za sterami", Href: "/pl/crew", Key: "crew"},
		{Label: "Nawiąż łączność", Href: "/pl/contact", Key: "contact"},
	},

	Strings: map[string]string{
		"nav.toggle_open":         "Otwórz menu",
		"nav.toggle_close":        "Zamknij menu",
		"nav.lang_label":          "Wybierz język",
		"footer.about":            "Polskie systemy bezzałogowe projektowane i produkowane lokalnie. Inżynieria, która działa tam, gdzie zawodzą standardowe rozwiązania.",
		"footer.nav_title":        "Nawigacja",
		"footer.contact_title":    "Kontakt",
		"footer.copyright":        "© Eagle Aerospace · Wszelkie prawa zastrzeżone",
		"flash.contact_success":   "Dziękujemy. Wiadomość trafiła do naszego zespołu - odezwiemy się tak szybko, jak to możliwe.",
		"flash.contact_error":     "Sprawdź podświetlone pola i spróbuj ponownie.",
		"err.contact.name":        "Podaj imię i nazwisko.",
		"err.contact.email_req":   "Podaj adres e-mail.",
		"err.contact.email_inv":   "Podaj poprawny adres e-mail.",
		"err.contact.topic":       "Wybierz temat zgłoszenia.",
		"err.contact.message_req": "Napisz krótką wiadomość.",
		"err.contact.message_min": "Wiadomość jest zbyt krótka - opisz pokrótce, czego dotyczy.",
	},

	Meta: map[string]PageMeta{
		"home": {
			Title:    "Eagle Aerospace - Systemy Bezzałogowe",
			MetaDesc: "Polska firma projektująca bezzałogowe statki powietrzne dla sektora publicznego, służb mundurowych i instytucji badawczych.",
		},
		"hangar": {
			Title:    "Hangar - Eagle Aerospace",
			MetaDesc: "Katalog platform bezzałogowych Eagle Aerospace - płatowce, VTOL, wielowirnikowce.",
		},
		"scenarios": {
			Title:    "Scenariusze - Eagle Aerospace",
			MetaDesc: "Inspekcje techniczne, rolnictwo precyzyjne, ratownictwo, monitoring środowiska - operacyjne zastosowania platform Eagle Aerospace.",
		},
		"cooperation": {
			Title:    "Współpraca - Eagle Aerospace",
			MetaDesc: "Współpraca Eagle Aerospace z uczelniami, jednostkami badawczymi i instytucjami publicznymi.",
		},
		"crew": {
			Title:    "Za sterami - Eagle Aerospace",
			MetaDesc: "Zespół inżynierów i konstruktorów Eagle Aerospace.",
		},
		"contact": {
			Title:    "Nawiąż łączność - Eagle Aerospace",
			MetaDesc: "Skontaktuj się z zespołem inżynierskim Eagle Aerospace.",
		},
	},

	Home: HomePageData{
		Tagline: "Eagle Aerospace · Rzeszów",
		Heading: "Bezzałogowe systemy lotnicze dla wymagających środowisk.",
		Lead: "Eagle Aerospace projektuje i produkuje bezzałogowe statki powietrzne dla " +
			"sektora publicznego, służb mundurowych oraz instytucji badawczych. Tworzymy " +
			"systemy, które działają tam, gdzie zawodzą standardowe rozwiązania: w trudnym " +
			"terenie i w wymagających warunkach.",
		CTACatalog: "Zobacz katalog platform",
		CTAContact: "Nawiąż łączność",
		Stats: []HomeStat{
			{Key: "Klasy platform", Value: "Płatowce · VTOL · Wielowirnikowce"},
			{Key: "Projektowanie", Value: "Lokalne, od podstaw"},
			{Key: "Lokalizacja", Value: "Rzeszów · Dolina Lotnicza"},
		},
		AreasEyebrow: "Obszary zastosowań",
		AreasHeading: "Co projektujemy i dla kogo",
		AreasLead:    "Trzy główne kierunki, w których nasze platformy są dziś wykorzystywane przez partnerów przemysłowych, akademickich i publicznych.",
		Highlights: []MediaCard{
			{
				Title:   "Inspekcje techniczne",
				Text:    "Przegląd infrastruktury liniowej, mostów, elektrowni i instalacji przemysłowych - bez angażowania ekip wspinaczkowych i bez przerywania pracy obiektu.",
				ImgPath: "/content/img/drone_scan.png",
				Alt:     "Drona inspekcjonująca instalację przemysłową",
			},
			{
				Title:   "Rozszerzone operacje",
				Text:    "Współpraca wielu platform w jednej misji: stacja naziemna, łącze danych, koordynacja i wymiana telemetrii w czasie rzeczywistym.",
				ImgPath: "/content/img/map_network.png",
				Alt:     "Sieć współpracujących platform nad obszarem operacji",
			},
			{
				Title:   "Mapowanie i rozpoznanie",
				Text:    "Wysokiej rozdzielczości skanowanie terenu, fotogrametria i ortofotomapy. Dane gotowe do dalszej analizy w GIS i systemach decyzyjnych.",
				ImgPath: "/content/img/scan_terrain.png",
				Alt:     "Mapowanie terenu i sieć transmisji danych",
			},
		},
	},

	Hangar: HangarPageData{
		Heading: "Hangar",
		Tagline: "Budujemy maszyny, które mają zadanie do wykonania",
		Lead: "Oferujemy szerokie spektrum platform bezzałogowych dostosowanych do różnych klas " +
			"zadań i środowisk operacyjnych. Katalog to punkt wyjścia - jeśli potrzebujesz " +
			"czegoś spoza listy, rozpoczniemy nowy projekt w oparciu o Twoje wymagania.",
		Groups: []CatalogGroup{
			{
				Label: "Klasa I - Płatowce",
				Items: []CatalogItem{
					{Name: "EA-F1", Description: "Zwrotne latające skrzydło do szybkiej reakcji na wykryte zagrożenie."},
					{Name: "EA-F2", Description: "Klasyczna platforma stałopłatowa do długotrwałego monitoringu przestrzeni i misji ISR."},
					{Name: "EA-F3", Description: "Kompaktowe latające skrzydło składane do rozmiaru plecaka. Przeznaczone do monitoringu i patrolowania terenu."},
				},
			},
			{
				Label: "Klasa II - VTOL",
				Items: []CatalogItem{
					// {Name: "EA-V1", Description: "VTOL hybrydowy - start pionowy, przelot poziomy, bez pasa startowego."},
					// {Name: "EA-V2", Description: "Tilt-rotor przeznaczony do długotrwałych misji patrolowych."},
					{Name: "EA-V3", Description: "Zmodyfikowana platforma F3 z wzmocnionym napędem umożliwiającym pionowy start."},
				},
			},
			{
				Label: "Klasa III - Wielowirnikowce",
				Items: []CatalogItem{
					{Name: "EA-M1", Description: "Quadcopter o rozpiętości 65 cm przystosowany do przenoszenia ładunku do 2 kg."},
					{Name: "EA-M2", Description: "Zwrotny quadcopter FPV do szybkich operacji w trudno dostępnym terenie."},
				//	{Name: "EA-M3", Description: "Mikro wielowirnikowiec do operacji wewnątrz budynków."},
				},
			},
		},
	},

	Scenarios: ScenariosPageData{
		Heading: "Scenariusze",
		Tagline: "Nie sprzedajemy dronów. Rozwiązujemy konkretne problemy.",
		Lead: "Systemy bezzałogowe Eagle Aerospace projektowane są z myślą o operacjach, " +
			"w których kluczowe znaczenie ma precyzja, zasięg oraz bezpieczeństwo ludzi. " +
			"Wspierają człowieka w środowiskach o podwyższonym stopniu ryzyka.",
		Scenarios: []ScenarioCard{
			{
				Title: "Inspekcje techniczne",
				Text:  "Przegląd infrastruktury liniowej, mostów, elektrowni i instalacji przemysłowych bez konieczności angażowania specjalistycznych ekip wspinaczkowych.",
			},
			{
				Title: "Rolnictwo precyzyjne",
				Text:  "Multispektralne mapowanie upraw, monitoring kondycji roślin i wsparcie opryskowe w oparciu o precyzyjne dane z lotu.",
			},
			{
				Title: "Ratownictwo",
				Text:  "Poszukiwania z powietrza, dostarczanie sprzętu do niedostępnych stref oraz koordynacja działań ratunkowych w trudnym terenie.",
			},
			{
				Title: "Monitoring środowiska",
				Text:  "Obserwacja obszarów chronionych, kontrola jakości powietrza, wykrywanie zanieczyszczeń i zmian w krajobrazie.",
			},
		},
	},

	Cooperation: CooperationPageData{
		Heading: "Współpraca",
		Tagline: "Przemysł i nauka we wspólnej przestrzeni powietrznej",
		Lead: "Eagle Aerospace rozwija się w ścisłej współpracy ze środowiskiem akademickim " +
			"i instytucjami badawczymi. Partnerstwo z Politechniką Rzeszowską, jednym " +
			"z wiodących ośrodków lotniczych w Polsce, pozwala nam łączyć najnowsze wyniki " +
			"badań z praktycznym projektowaniem systemów bezzałogowych.",
		LeadImgPath:   "/content/img/map_network.png",
		LeadImgAlt:    "Sieć współpracujących ośrodków badawczych w Polsce",
		PartnersTitle: "Partnerzy i instytucje",
		Media: []MediaCard{
			{
				Title: "Wspólne projekty badawcze",
				Text:  "Konstrukcja płatowców, charakterystyki aerodynamiczne i nowe rozwiązania w obszarze napędów elektrycznych.",
			},
			{
				Title: "Algorytmy analizy obrazu",
				Text:  "Rozwój algorytmów opartych na uczeniu maszynowym do automatycznej detekcji obiektów i klasyfikacji terenu.",
			},
		},
		Partners: []PartnerCard{
			{Name: "Politechnika Rzeszowska"},
			{Name: "Dolina Lotnicza"},
			{Name: "Instytut Lotnictwa"},
			{Name: "Partner badawczy"},
			{Name: "Partner przemysłowy"},
			{Name: "Partner publiczny"},
		},
	},

	Crew: CrewPageData{
		Heading: "Za sterami",
		Tagline: "Pasja to za małe słowo. Nazywamy to obsesją na punkcie jakości.",
		Lead: "Eagle Aerospace to zespół inżynierów, konstruktorów i specjalistów, dla których " +
			"bezzałogowe systemy powietrzne to nie branża - to pasja przekuta w zawód. " +
			"Jako startup zachowujemy pełną elastyczność i zaangażowanie. Każdy projekt jest " +
			"dla nas ważny, każdy klient ma bezpośredni kontakt z zespołem.",
		Team: []TeamMember{
			{Name: "Grzegorz Trzęsień", Role: "Aviation Systems Engineer", Bio: "Mechanika lotu, planowanie misji, koncepcja systemu"},
			{Name: "Filip Starzyński", Role: "Aviation Systems Engineer", Bio: "Specjalista od aerodynamiki i analiz numerycznych, obróbka CNC, prototypowanie"},
			{Name: "Aleksander Słomka", Role: "Software Engineer", Bio: "Systemy pokładowe, łącza danych i autopiloty."},
			{Name: "Jędrzej Srebrzyński", Role: "Aviation Systems Engineer", Bio: "Koncepcja systemu, integracja komponentów"},
			{Name: "Kajetan Ostrowski", Role: "Aviation Systems Engineer", Bio: "Projektowanie CAD, prototypowanie, relacje z partnerami"},
		},
	},

	Contact: ContactPageData{
		Heading: "Nawiąż łączność",
		Tagline: "Bezpośredni kontakt z zespołem inżynierskim",
		Lead: "Masz pytanie o konkretną platformę, chcesz omówić wymagania projektu lub po " +
			"prostu sprawdzić, czy jesteśmy właściwym partnerem - napisz lub zadzwoń. " +
			"Rozmawiasz bezpośrednio z inżynierami. Znajdziesz nas w Rzeszowie, " +
			"w sercu polskiego przemysłu lotniczego.",
		Topics: []string{
			"Zapytanie o konkretną platformę",
			"Nowy projekt",
			"Współpraca badawcza",
			"Inne",
		},
		ContactInfo: ContactInfoBlock{
			Email:   "biuro@eagleaero.space",
			Phone:   "+48 883 550 701",
			Address: []string{"ul. Słowackiego 7A/7 35-060 Rzeszów, Polska"},
			KRS: "0001212701",
			NIP: "8172223670",
			REGON: "543537700",
		},
		Labels: ContactLabels{
			HoneypotLabel:    "Strona",
			Name:             "Imię i nazwisko",
			Org:              "Instytucja / Organizacja",
			Email:            "Adres e-mail",
			Phone:            "Telefon",
			Topic:            "Temat",
			Message:          "Wiadomość",
			TopicPlaceholder: "Wybierz temat",
			Submit:           "Wyślij wiadomość",
			AsideTitle:       "Kontakt bezpośredni",
			EmailLabel:       "E-mail",
			PhoneLabel:       "Telefon",
			AddressLabel:     "Adres",
			CompanyLabel:     "Dane firmy",
		},
	},
}
