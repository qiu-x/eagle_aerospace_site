package submission

import (
	"fmt"
	"log/slog"
	"time"
)

type Submission struct {
	ReceivedAt time.Time `json:"received_at"`
	Lang       string    `json:"lang"`
	Name       string    `json:"name"`
	Org        string    `json:"org,omitempty"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone,omitempty"`
	Topic      string    `json:"topic"`
	Message    string    `json:"message"`
}

type Sink struct {
	SMTP     *SMTPConfig
	FilePath string
	Logger   *slog.Logger
}

func New(smtp *SMTPConfig, filePath string, logger *slog.Logger) *Sink {
	if logger == nil {
		logger = slog.Default()
	}
	if smtp != nil && smtp.Addr == "" {
		smtp = nil
	}
	return &Sink{SMTP: smtp, FilePath: filePath, Logger: logger}
}

func (s *Sink) Configured() bool {
	return s.SMTP != nil || s.FilePath != ""
}

func (s *Sink) Submit(sub Submission) error {
	if sub.ReceivedAt.IsZero() {
		sub.ReceivedAt = time.Now().UTC()
	}

	if s.FilePath == "" && s.SMTP == nil {
		return fmt.Errorf("no submission backend configured")
	}

	var fileErr error
	if s.FilePath != "" {
		fileErr = appendFile(s.FilePath, sub)
		if fileErr == nil {
			s.Logger.Info("submission written to file",
				slog.String("file", s.FilePath),
				slog.String("email", sub.Email))
		} else {
			s.Logger.Error("file write failed",
				slog.String("file", s.FilePath),
				slog.String("email", sub.Email),
				slog.Any("err", fileErr))
		}
	}

	var smtpErr error
	if s.SMTP != nil {
		smtpErr = sendSMTP(s.SMTP, sub)
		if smtpErr == nil {
			s.Logger.Info("submission sent via smtp",
				slog.String("email", sub.Email),
				slog.String("topic", sub.Topic))
		} else {
			s.Logger.Error("smtp send failed",
				slog.String("email", sub.Email),
				slog.Any("err", smtpErr))
		}
	}

	// Only fail when every configured backend errored.
	switch {
	case s.FilePath != "" && fileErr == nil:
		return nil
	case s.SMTP != nil && smtpErr == nil:
		return nil
	case fileErr != nil && smtpErr != nil:
		return fmt.Errorf("file backup failed (%w) and smtp send failed: %w", fileErr, smtpErr)
	case fileErr != nil:
		return fmt.Errorf("file write failed: %w", fileErr)
	default:
		return fmt.Errorf("smtp send failed: %w", smtpErr)
	}
}
