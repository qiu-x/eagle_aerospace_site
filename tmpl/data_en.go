package tmpl

var siteEN = SiteData{
	NavItems: []NavItem{
		{Label: "Hangar", Href: "/en/hangar", Key: "hangar"},
		{Label: "Scenarios", Href: "/en/scenarios", Key: "scenarios"},
		// {Label: "Cooperation", Href: "/en/cooperation", Key: "cooperation"},
		{Label: "Crew", Href: "/en/crew", Key: "crew"},
		{Label: "Get in touch", Href: "/en/contact", Key: "contact"},
	},

	Strings: map[string]string{
		"nav.toggle_open":         "Open menu",
		"nav.toggle_close":        "Close menu",
		"nav.lang_label":          "Choose language",
		"footer.about":            "Polish unmanned aerial systems, designed and built locally. Engineering that performs where standard solutions fall short.",
		"footer.nav_title":        "Navigation",
		"footer.contact_title":    "Contact",
		"footer.copyright":        "© Eagle Aerospace · All rights reserved",
		"flash.contact_success":   "Thank you. Your message reached our team - we'll get back to you as soon as possible.",
		"flash.contact_error":     "Please review the highlighted fields and try again.",
		"err.contact.name":        "Please provide your full name.",
		"err.contact.email_req":   "Please provide an email address.",
		"err.contact.email_inv":   "Please provide a valid email address.",
		"err.contact.topic":       "Choose a subject for your enquiry.",
		"err.contact.message_req": "Please write a brief message.",
		"err.contact.message_min": "The message is too short - please outline what it concerns.",
	},

	Meta: map[string]PageMeta{
		"home": {
			Title:    "Eagle Aerospace - Unmanned Aerial Systems",
			MetaDesc: "Polish company designing unmanned aerial vehicles for the public sector, uniformed services and research institutions.",
		},
		"hangar": {
			Title:    "Hangar - Eagle Aerospace",
			MetaDesc: "Eagle Aerospace catalogue of unmanned platforms - fixed-wing, VTOL and multirotor aircraft.",
		},
		"scenarios": {
			Title:    "Scenarios - Eagle Aerospace",
			MetaDesc: "Technical inspections, precision agriculture, search and rescue, environmental monitoring - operational use cases for Eagle Aerospace platforms.",
		},
		"cooperation": {
			Title:    "Cooperation - Eagle Aerospace",
			MetaDesc: "Eagle Aerospace cooperation with universities, research units and public institutions.",
		},
		"crew": {
			Title:    "Crew - Eagle Aerospace",
			MetaDesc: "The engineering and design team behind Eagle Aerospace.",
		},
		"contact": {
			Title:    "Get in touch - Eagle Aerospace",
			MetaDesc: "Reach out to the Eagle Aerospace engineering team.",
		},
	},

	Home: HomePageData{
		Tagline: "Eagle Aerospace · Rzeszów",
		Heading: "Unmanned aerial systems for demanding environments.",
		Lead: "Eagle Aerospace designs and manufactures unmanned aerial vehicles for the " +
			"public sector, uniformed services and research institutions. We build systems " +
			"that perform where standard solutions fall short - in difficult terrain and " +
			"under demanding conditions.",
		CTACatalog: "Browse the platform catalogue",
		CTAContact: "Get in touch",
		Stats: []HomeStat{
			{Key: "Platform classes", Value: "Fixed-wing · VTOL · Multirotor"},
			{Key: "Engineering", Value: "Local, from the ground up"},
			{Key: "Location", Value: "Rzeszów · Aviation Valley"},
		},
		AreasEyebrow: "Application areas",
		AreasHeading: "What we build, and who it's for",
		AreasLead:    "Three primary directions in which our platforms are used today by industrial, academic and public-sector partners.",
		Highlights: []MediaCard{
			{
				Title:   "Technical inspections",
				Text:    "Survey of linear infrastructure, bridges, power plants and industrial installations - without rope-access teams and without halting operations.",
				ImgPath: "/content/img/drone_scan.png",
				Alt:     "Drone inspecting an industrial installation",
			},
			{
				Title:   "Extended operations",
				Text:    "Multi-platform cooperation in a single mission: ground station, data link, coordination and real-time telemetry exchange.",
				ImgPath: "/content/img/map_network.png",
				Alt:     "Network of cooperating platforms over an operational area",
			},
			{
				Title:   "Mapping and reconnaissance",
				Text:    "High-resolution terrain scanning, photogrammetry and orthophoto maps. Data ready for downstream analysis in GIS and decision-support systems.",
				ImgPath: "/content/img/scan_terrain.png",
				Alt:     "Terrain mapping and data transmission network",
			},
		},
	},

	Hangar: HangarPageData{
		Heading: "Hangar",
		Tagline: "We build aircraft with a job to do",
		Lead: "We offer a broad spectrum of unmanned platforms tailored to different mission " +
			"classes and operating environments. The catalogue is a starting point - if you " +
			"need something outside this list, we'll start a new project around your requirements.",
		Groups: []CatalogGroup{
			{
				Label: "Class I - Fixed-wing",
				Items: []CatalogItem{
					{Name: "EA-F1", Description: "Agile flying wing for rapid response to detected threats."},
					{Name: "EA-F2", Description: "Fixed-wing platform for long-endurance area monitoring and ISR missions."},
					{Name: "EA-F3", Description: "Compact flying wing that folds down to backpack size. Designed for monitoring and area patrol."},
				},
			},
			{
				Label: "Class II - VTOL",
				Items: []CatalogItem{
					// {Name: "EA-V1", Description: "Hybrid VTOL - vertical take-off, horizontal cruise, no runway required."},
					// {Name: "EA-V2", Description: "Tilt-rotor designed for long-endurance patrol missions."},
					{Name: "EA-V3", Description: "Modified F3 platform with reinforced propulsion enabling vertical take-off and landing."},
				},
			},
			{
				Label: "Class III - Multirotor",
				Items: []CatalogItem{
					 {Name: "EA-M1", Description: "Quadcopter with a 65 cm diagonal, rated for payloads up to 2 kg."},
					 {Name: "EA-M2", Description: "Agile FPV quadcopter for fast operations in hard-to-reach terrain."},
			//		{Name: "EA-M3", Description: "Micro multirotor for indoor operations."},
				},
			},
		},
	},

	Scenarios: ScenariosPageData{
		Heading: "Scenarios",
		Tagline: "We don't sell drones. We solve specific problems.",
		Lead: "Eagle Aerospace unmanned systems are designed for operations where precision, " +
			"range and human safety are critical. They support people in environments with " +
			"elevated risk.",
		Scenarios: []ScenarioCard{
			{
				Title: "Technical inspections",
				Text:  "Survey of linear infrastructure, bridges, power plants and industrial installations without engaging specialised rope-access teams.",
			},
			{
				Title: "Precision agriculture",
				Text:  "Multispectral crop mapping, plant-health monitoring and spraying support based on precise in-flight data.",
			},
			{
				Title: "Search and rescue",
				Text:  "Aerial search, equipment delivery into inaccessible zones and coordination of rescue operations in difficult terrain.",
			},
			{
				Title: "Environmental monitoring",
				Text:  "Observation of protected areas, air-quality control, detection of pollution and changes in the landscape.",
			},
		},
	},

	Cooperation: CooperationPageData{
		Heading: "Cooperation",
		Tagline: "Industry and science sharing the same airspace",
		Lead: "Eagle Aerospace develops in close cooperation with the academic community and " +
			"research institutions. Our partnership with the Rzeszów University of Technology, " +
			"one of Poland's leading aviation centres, lets us combine the latest research " +
			"with hands-on design of unmanned systems.",
		LeadImgPath:   "/content/img/map_network.png",
		LeadImgAlt:    "Network of cooperating research centres in Poland",
		PartnersTitle: "Partners and institutions",
		Media: []MediaCard{
			{
				Title: "Joint research projects",
				Text:  "Airframe construction, aerodynamic characteristics and new approaches to electric propulsion.",
			},
			{
				Title: "Image-analysis algorithms",
				Text:  "Machine-learning-based algorithms for automatic object detection and terrain classification.",
			},
		},
		Partners: []PartnerCard{
			{Name: "Rzeszów University of Technology"},
			{Name: "Aviation Valley"},
			{Name: "Institute of Aviation"},
			{Name: "Research partner"},
			{Name: "Industrial partner"},
			{Name: "Public-sector partner"},
		},
	},

	Crew: CrewPageData{
		Heading: "Crew",
		Tagline: "Passion is too small a word. We call it an obsession with quality.",
		Lead: "Eagle Aerospace is a team of engineers, designers and specialists for whom " +
			"unmanned aerial systems are not an industry - they are a passion turned into " +
			"a profession. As a startup we keep full flexibility and commitment. Every " +
			"project matters to us; every client talks to the team directly.",
		Team: []TeamMember{
			{Name: "Grzegorz Trzęsień", Role: "Aviation Systems Engineer", Bio: "Flight mechanics, mission planning, system concept."},
			{Name: "Filip Starzyński", Role: "Aviation Systems Engineer", Bio: "Aerodynamics and numerical analysis, CNC machining, prototyping."},
			{Name: "Aleksander Słomka", Role: "Software Engineer", Bio: "On-board systems, data links and autopilots."},
			{Name: "Jędrzej Srebrzyński", Role: "Aviation Systems Engineer", Bio: "System concept, component integration."},
			{Name: "Kajetan Ostrowski", Role: "Aviation Systems Engineer", Bio: "CAD design, prototyping, partner relations."},
		},
	},

	Contact: ContactPageData{
		Heading: "Get in touch",
		Tagline: "Direct contact with the engineering team",
		Lead: "Have a question about a specific platform, want to discuss project requirements " +
			"or simply check whether we're the right partner - write or call. You'll be " +
			"talking directly to engineers. You'll find us in Rzeszów, at the heart of the " +
			"Polish aerospace industry.",
		Topics: []string{
			"Enquiry about a specific platform",
			"New project",
			"Research cooperation",
			"Other",
		},
		ContactInfo: ContactInfoBlock{
			Email:   "biuro@eagleaero.space",
			Phone:   "+48 883 550 701",
			Address: []string{"ul. Słowackiego 7A/7 35-060 Rzeszów, Poland"},
			KRS: "0001212701",
			NIP: "8172223670",
			REGON: "543537700",
		},
		Labels: ContactLabels{
			HoneypotLabel:    "Website",
			Name:             "Full name",
			Org:              "Institution / Organisation",
			Email:            "Email address",
			Phone:            "Phone",
			Topic:            "Subject",
			Message:          "Message",
			TopicPlaceholder: "Choose a subject",
			Submit:           "Send message",
			AsideTitle:       "Direct contact",
			EmailLabel:       "Email",
			PhoneLabel:       "Phone",
			AddressLabel:     "Address",
			CompanyLabel:     "Company info",
		},
	},
}
