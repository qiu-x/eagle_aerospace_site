package tmpl

// Single primary-navigation entry.
type NavItem struct {
	Label string
	Href  string
	Key   string
}

// Shown at the top of a page (e.g. contact-form success).
type FlashMessage struct {
	Kind string // "success" | "error"
	Text string
}

// Shell payload every page receives.
type BasePageData struct {
	Title        string
	MetaDesc     string
	ActiveNavKey string
	BodyClass    string
	NavItems     []NavItem
	Flash        *FlashMessage
	Content      any
}

/* 
 * Reusable component structs:
 */

type MediaCard struct {
	Title   string
	Text    string
	ImgPath string
	Alt     string
}

type CatalogItem struct {
	Name        string
	Description string
	ImgPath     string
	Alt         string
}

type CatalogGroup struct {
	Label string
	Items []CatalogItem
}

type ScenarioCard struct {
	Title   string
	Text    string
	ImgPath string
	Alt     string
}

type PartnerCard struct {
	Name     string
	LogoPath string
	Link     string
}

type TeamMember struct {
	Name    string
	Role    string
	Bio     string
	ImgPath string
	Alt     string
}

type ContactInfoBlock struct {
	Email   string
	Phone   string
	Address []string // address lines
}

/* 
 * Page-specific payloads: 
 */

type HomeStat struct {
	Key   string
	Value string
}

type HomePageData struct {
	Tagline    string
	Lead       string
	Stats      []HomeStat
	Highlights []MediaCard
}

type HangarPageData struct {
	Tagline string
	Lead    string
	Groups  []CatalogGroup
}

type ScenariosPageData struct {
	Tagline   string
	Lead      string
	Scenarios []ScenarioCard
}

type CooperationPageData struct {
	Tagline     string
	Lead        string
	LeadImgPath string
	LeadImgAlt  string
	Media       []MediaCard
	Partners    []PartnerCard
}

type CrewPageData struct {
	Tagline string
	Lead    string
	Team    []TeamMember
}

// User-submitted contact form payload.
type ContactFormInput struct {
	Name    string
	Org     string
	Email   string
	Phone   string
	Topic   string
	Message string
}

type ContactPageData struct {
	Tagline     string
	Lead        string
	Topics      []string
	Form        ContactFormInput
	FieldErrors map[string]string
	ContactInfo ContactInfoBlock
}
