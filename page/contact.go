package page

import (
	"log/slog"
	"net/http"
	"net/mail"
	"strings"

	"eagle_site_v2/tmpl"
)

const minMessageLen = 20

func contactBaseData(content tmpl.ContactPageData, flash *tmpl.FlashMessage) tmpl.BasePageData {
	return tmpl.BasePageData{
		Title:        "Nawiąż łączność - Eagle Aerospace",
		MetaDesc:     "Skontaktuj się z zespołem inżynierskim Eagle Aerospace.",
		ActiveNavKey: "contact",
		BodyClass:    "page-contact",
		NavItems:     tmpl.NavItems,
		Flash:        flash,
		Content:      content,
	}
}

func ContactGetHandler(r *tmpl.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var flash *tmpl.FlashMessage
		if req.URL.Query().Get("sent") == "1" {
			flash = &tmpl.FlashMessage{
				Kind: "success",
				Text: "Dziękujemy. Wiadomość trafiła do naszego zespołu - odezwiemy się tak szybko, jak to możliwe.",
			}
		}
		// Copy the static content so we can attach a per-request blank form.
		c := tmpl.ContactContent
		c.Form = tmpl.ContactFormInput{}
		c.FieldErrors = nil
		renderOrFail(w, r, "contact", contactBaseData(c, flash))
	}
}

func ContactPostHandler(r *tmpl.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil {
			http.Error(w, "bad form", http.StatusBadRequest)
			return
		}

		// Honeypot :D
		if strings.TrimSpace(req.PostForm.Get("website")) != "" {
			http.Redirect(w, req, "/contact?sent=1", http.StatusSeeOther)
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

		errs := validateContact(form)
		if len(errs) > 0 {
			c := tmpl.ContactContent
			c.Form = form
			c.FieldErrors = errs
			flash := &tmpl.FlashMessage{
				Kind: "error",
				Text: "Sprawdź podświetlone pola i spróbuj ponownie.",
			}
			w.WriteHeader(http.StatusUnprocessableEntity)
			renderOrFail(w, r, "contact", contactBaseData(c, flash))
			return
		}

		// v1: just log it. SMTP / file storage comes later.
		slog.Info("contact submission",
			slog.String("name", form.Name),
			slog.String("org", form.Org),
			slog.String("email", form.Email),
			slog.String("phone", form.Phone),
			slog.String("topic", form.Topic),
			slog.Int("message_len", len(form.Message)),
		)

		http.Redirect(w, req, "/contact?sent=1", http.StatusSeeOther)
	}
}

func validateContact(f tmpl.ContactFormInput) map[string]string {
	errs := map[string]string{}

	if f.Name == "" {
		errs["name"] = "Podaj imię i nazwisko."
	}
	if f.Email == "" {
		errs["email"] = "Podaj adres e-mail."
	} else if addr, err := mail.ParseAddress(f.Email); err != nil || addr.Address != f.Email {
		errs["email"] = "Podaj poprawny adres e-mail."
	}
	if f.Topic == "" {
		errs["topic"] = "Wybierz temat zgłoszenia."
	}
	if f.Message == "" {
		errs["message"] = "Napisz krótką wiadomość."
	} else if len(f.Message) < minMessageLen {
		errs["message"] = "Wiadomość jest zbyt krótka - opisz pokrótce, czego dotyczy."
	}

	return errs
}
