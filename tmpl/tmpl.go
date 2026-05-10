package tmpl

type Lang string

const (
	LangPL Lang = "pl"
	LangEN Lang = "en"
)

// Navigation entry.
type NavItem struct {
	Label string
	Href  string
	Key   string
}

// Language toggle entry (in nav-bar).
type LangLink struct {
	Code   Lang
	Label  string
	Href   string
	Active bool
}

// Per-page <title> + <meta name="description">.
type PageMeta struct {
	Title    string
	MetaDesc string
}

// Shown at the top of a page (e.g. contact-form success).
type FlashMessage struct {
	Kind string // "success" | "error"
	Text string
}

// Shell payload every page receives.
type BasePageData struct {
	Lang         Lang
	Title        string
	MetaDesc     string
	ActiveNavKey string
	BodyClass    string
	HomeHref     string
	NavItems     []NavItem
	Strings      map[string]string
	LangLinks    []LangLink
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
	Tagline      string
	Heading      string
	Lead         string
	CTACatalog   string
	CTAContact   string
	Stats        []HomeStat
	AreasEyebrow string
	AreasHeading string
	AreasLead    string
	Highlights   []MediaCard
}

type HangarPageData struct {
	Heading string
	Tagline string
	Lead    string
	Groups  []CatalogGroup
}

type ScenariosPageData struct {
	Heading   string
	Tagline   string
	Lead      string
	Scenarios []ScenarioCard
}

type CooperationPageData struct {
	Heading       string
	Tagline       string
	Lead          string
	LeadImgPath   string
	LeadImgAlt    string
	Media         []MediaCard
	Partners      []PartnerCard
	PartnersTitle string
}

type CrewPageData struct {
	Heading string
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

// Localized chrome labels for the contact form.
type ContactLabels struct {
	HoneypotLabel    string
	Name             string
	Org              string
	Email            string
	Phone            string
	Topic            string
	Message          string
	TopicPlaceholder string
	Submit           string
	AsideTitle       string
	EmailLabel       string
	PhoneLabel       string
	AddressLabel     string
}

type ContactPageData struct {
	Heading     string
	Tagline     string
	Lead        string
	Topics      []string
	Form        ContactFormInput
	FieldErrors map[string]string
	ContactInfo ContactInfoBlock
	Labels      ContactLabels
}
