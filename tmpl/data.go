package tmpl

var NavItems = []NavItem{
	{Label: "Hangar", Href: "/hangar", Key: "hangar"},
	{Label: "Scenariusze", Href: "/scenarios", Key: "scenarios"},
	// {Label: "Współpraca", Href: "/cooperation", Key: "cooperation"},
	{Label: "Za sterami", Href: "/crew", Key: "crew"},
	{Label: "Nawiąż łączność", Href: "/contact", Key: "contact"},
}

var HomeContent = HomePageData{
	Tagline: "Eagle Aerospace · Rzeszów",
	Lead: "Eagle Aerospace projektuje i produkuje bezzałogowe statki powietrzne dla " +
		"sektora publicznego, służb mundurowych oraz instytucji badawczych. Tworzymy " +
		"systemy, które działają tam, gdzie zawodzą standardowe rozwiązania: w trudnym " +
		"terenie i w wymagających warunkach.",
	Stats: []HomeStat{
		{Key: "Klasy platform", Value: "Płatowce · VTOL · Wielowirnikowce"},
		{Key: "Projektowanie", Value: "Lokalne, od podstaw"},
		{Key: "Lokalizacja", Value: "Rzeszów · Dolina Lotnicza"},
	},
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
}

var HangarContent = HangarPageData{
	Tagline: "Budujemy maszyny, które mają zadanie do wykonania",
	Lead: "Oferujemy szerokie spektrum platform bezzałogowych dostosowanych do różnych klas " +
		"zadań i środowisk operacyjnych. Katalog to punkt wyjścia - jeśli potrzebujesz " +
		"czegoś spoza listy, rozpoczniemy nowy projekt w oparciu o Twoje wymagania.",
	Groups: []CatalogGroup{
		{
			Label: "Klasa I - Płatowce",
			Items: []CatalogItem{
				{Name: "EA-F1", Description: "Stałopłat dalekiego zasięgu do rozpoznania i mapowania terenu."},
				{Name: "EA-F2", Description: "Kompaktowy płatowiec do szybkich misji zwiadowczych w trudnym terenie."},
				{Name: "EA-F3", Description: "Platforma o zwiększonym udźwigu dla zaawansowanego osprzętu pomiarowego."},
			},
		},
		{
			Label: "Klasa II - VTOL",
			Items: []CatalogItem{
				{Name: "EA-V1", Description: "VTOL hybrydowy - start pionowy, przelot poziomy, bez pasa startowego."},
				{Name: "EA-V2", Description: "Tilt-rotor przeznaczony do długotrwałych misji patrolowych."},
				{Name: "EA-V3", Description: "VTOL o skróconym przebiegu przeznaczony dla jednostek mobilnych."},
			},
		},
		{
			Label: "Klasa III - Wielowirnikowce",
			Items: []CatalogItem{
				{Name: "EA-M1", Description: "Hexakopter inspekcyjny z precyzyjnym systemem zawisu."},
				{Name: "EA-M2", Description: "Oktorotor do misji w środowiskach przemysłowych i infrastrukturalnych."},
				{Name: "EA-M3", Description: "Mikro wielowirnikowiec do operacji wewnątrz budynków."},
			},
		},
	},
}

var ScenariosContent = ScenariosPageData{
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
}

var CooperationContent = CooperationPageData{
	Tagline: "Przemysł i nauka we wspólnej przestrzeni powietrznej",
	Lead: "Eagle Aerospace rozwija się w ścisłej współpracy ze środowiskiem akademickim " +
		"i instytucjami badawczymi. Partnerstwo z Politechniką Rzeszowską, jednym " +
		"z wiodących ośrodków lotniczych w Polsce, pozwala nam łączyć najnowsze wyniki " +
		"badań z praktycznym projektowaniem systemów bezzałogowych.",
	LeadImgPath: "/content/img/map_network.png",
	LeadImgAlt:  "Sieć współpracujących ośrodków badawczych w Polsce",
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
}

var CrewContent = CrewPageData{
	Tagline: "Pasja to za małe słowo. Nazywamy to obsesją na punkcie jakości.",
	Lead: "Eagle Aerospace to zespół inżynierów, konstruktorów i specjalistów, dla których " +
		"bezzałogowe systemy powietrzne to nie branża - to pasja przekuta w zawód. " +
		"Jako startup zachowujemy pełną elastyczność i zaangażowanie. Każdy projekt jest " +
		"dla nas ważny, każdy klient ma bezpośredni kontakt z zespołem.",
	Team: []TeamMember{
		{Name: "Grzegorz Trzęsień", Role: "Aviation Systems Engineer", Bio: "Specjalista od aerodynamiki i materiałów kompozytowych."},
		{Name: "Filip Starzyński", Role: "Aviation Systems Engineer", Bio: "Specjalista od aerodynamiki i materiałów kompozytowych."},
		{Name: "Aleksander Słomka", Role: "Software Engineer", Bio: "Systemy pokładowe, łącza danych i autopiloty."},
		{Name: "Jędrzej Srebrzyński", Role: "Aviation Systems Engineer", Bio: "Specjalista od aerodynamiki i materiałów kompozytowych."},
		{Name: "Kajetan Ostrowski", Role: "AI Engineer", Bio: "Rozwój w zakresie dużych modelów językowych"},
	},
}

var ContactContent = ContactPageData{
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
		Address: []string{"Rzeszów, Polska", "Dolina Lotnicza"},
	},
}
