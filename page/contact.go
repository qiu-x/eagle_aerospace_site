package page

import (
	"log/slog"
	"net/http"
	"net/mail"
	"strings"

	"eagle_site_v2/tmpl"
)

const minMessageLen = 20

func contactBaseData(lang tmpl.Lang, reqPath string,
	content tmpl.ContactPageData, flash *tmpl.FlashMessage) tmpl.BasePageData {
	return newBaseData(lang, "contact", "page-contact", "contact", reqPath, content, flash)
}

func ContactGetHandler(r *tmpl.Renderer, lang tmpl.Lang) http.HandlerFunc {
	site := tmpl.SiteFor(lang)
	return func(w http.ResponseWriter, req *http.Request) {
		var flash *tmpl.FlashMessage
		if req.URL.Query().Get("sent") == "1" {
			flash = &tmpl.FlashMessage{
				Kind: "success",
				Text: site.Strings["flash.contact_success"],
			}
		}
		// Copy the static content so we can attach a per-request blank form.
		c := site.Contact
		c.Form = tmpl.ContactFormInput{}
		c.FieldErrors = nil
		renderOrFail(w, r, "contact", contactBaseData(lang, req.URL.Path, c, flash))
	}
}

func ContactPostHandler(r *tmpl.Renderer, lang tmpl.Lang) http.HandlerFunc {
	site := tmpl.SiteFor(lang)
	return func(w http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil {
			http.Error(w, "bad form", http.StatusBadRequest)
			return
		}

		// Honeypot :D
		if strings.TrimSpace(req.PostForm.Get("website")) != "" {
			http.Redirect(w, req, req.URL.Path+"?sent=1", http.StatusSeeOther)
			return
		}

		form := tmpl.ContactFormInput{
			Name:    strings.TrimSpace(req.PostForm.Get("name")),
			Org:     strings.TrimSpace(req.PostForm.Get("org")),
			Email:   strings.TrimSpace(req.PostForm.Get("email")),
			Phone:   strings.TrimSpace(req.PostForm.Get("phone")),
			Topic:   strings.TrimSpace(req.PostForm.Get("topic")),
			Message: strings.TrimSpace(req.PostForm.Get("message")),
		}

		errs := validateContact(form, site.Strings)
		if len(errs) > 0 {
			c := site.Contact
			c.Form = form
			c.FieldErrors = errs
			flash := &tmpl.FlashMessage{
				Kind: "error",
				Text: site.Strings["flash.contact_error"],
			}
			w.WriteHeader(http.StatusUnprocessableEntity)
			renderOrFail(w, r, "contact", contactBaseData(lang, req.URL.Path, c, flash))
			return
		}

		// SMTP / file storage comes later.
		slog.Info("contact submission",
			slog.String("lang", string(lang)),
			slog.String("name", form.Name),
			slog.String("org", form.Org),
			slog.String("email", form.Email),
			slog.String("phone", form.Phone),
			slog.String("topic", form.Topic),
			slog.Int("message_len", len(form.Message)),
		)

		http.Redirect(w, req, req.URL.Path+"?sent=1", http.StatusSeeOther)
	}
}

func validateContact(f tmpl.ContactFormInput, s map[string]string) map[string]string {
	errs := map[string]string{}

	if f.Name == "" {
		errs["name"] = s["err.contact.name"]
	}
	if f.Email == "" {
		errs["email"] = s["err.contact.email_req"]
	} else if addr, err := mail.ParseAddress(f.Email); err != nil || addr.Address != f.Email {
		errs["email"] = s["err.contact.email_inv"]
	}
	if f.Topic == "" {
		errs["topic"] = s["err.contact.topic"]
	}
	if f.Message == "" {
		errs["message"] = s["err.contact.message_req"]
	} else if len(f.Message) < minMessageLen {
		errs["message"] = s["err.contact.message_min"]
	}

	return errs
}
