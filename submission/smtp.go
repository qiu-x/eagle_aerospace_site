package submission

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

type SMTPConfig struct {
	Addr string
	User string
	Pass string
	From string
	To   string
}

func sendSMTP(c *SMTPConfig, sub Submission) error {
	if c.From == "" || c.To == "" {
		return fmt.Errorf("smtp: from and to are required")
	}

	var auth smtp.Auth
	if c.User != "" {
		host := c.Addr
		if i := strings.LastIndex(host, ":"); i > 0 {
			host = host[:i]
		}
		auth = smtp.PlainAuth("", c.User, c.Pass, host)
	}

	msg := buildMessage(c.From, c.To, sub)
	return smtp.SendMail(c.Addr, auth, c.From, []string{c.To}, []byte(msg))
}

func buildMessage(from, to string, sub Submission) string {
	var b strings.Builder
	subject := fmt.Sprintf("Eagle: contact (%s) - %s", sub.Lang, sub.Topic)
	encoded := "=?utf-8?B?" + base64.StdEncoding.EncodeToString([]byte(subject)) + "?="

	fmt.Fprintf(&b, "From: %s\r\n", from)
	fmt.Fprintf(&b, "To: %s\r\n", to)
	fmt.Fprintf(&b, "Reply-To: %s\r\n", sub.Email)
	fmt.Fprintf(&b, "Subject: %s\r\n", encoded)
	fmt.Fprintf(&b, "Date: %s\r\n", sub.ReceivedAt.Format(time.RFC1123Z))
	fmt.Fprintf(&b, "MIME-Version: 1.0\r\n")
	fmt.Fprintf(&b, "Content-Type: text/plain; charset=utf-8\r\n")
	fmt.Fprintf(&b, "Content-Transfer-Encoding: 8bit\r\n")
	fmt.Fprintf(&b, "\r\n")
	fmt.Fprintf(&b, "Eagle Aerospace - new contact submission\r\n\r\n")
	fmt.Fprintf(&b, "Lang:    %s\r\n", sub.Lang)
	fmt.Fprintf(&b, "Time:    %s\r\n", sub.ReceivedAt.UTC().Format(time.RFC3339))
	fmt.Fprintf(&b, "Name:    %s\r\n", sub.Name)
	fmt.Fprintf(&b, "Org:     %s\r\n", sub.Org)
	fmt.Fprintf(&b, "Email:   %s\r\n", sub.Email)
	fmt.Fprintf(&b, "Phone:   %s\r\n", sub.Phone)
	fmt.Fprintf(&b, "Topic:   %s\r\n\r\n", sub.Topic)
	fmt.Fprintf(&b, "Message:\r\n%s\r\n", sub.Message)
	return b.String()
}
