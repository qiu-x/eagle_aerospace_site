package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"eagle_site_v2/submission"
	"eagle_site_v2/tmpl"
)

type Config struct {
	Host string
	Port string

	SMTPAddr        string
	SMTPUser        string
	SMTPPass        string
	SMTPFrom        string
	SMTPTo          string
	SubmissionsFile string
}

func newConfig(args []string) (Config, error) {
	var cfg Config
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.StringVar(&cfg.Host, "host", "0.0.0.0", "Bind address")
	flags.StringVar(&cfg.Port, "port", "8080", "Listen port")

	flags.StringVar(&cfg.SMTPAddr, "smtp-addr", "",
		"SMTP server (host:port, e.g. smtp.example.com:587). Empty disables SMTP.")
	flags.StringVar(&cfg.SMTPUser, "smtp-user", "", "SMTP auth user (optional).")
	flags.StringVar(&cfg.SMTPPass, "smtp-pass", "", "SMTP auth password (optional).")
	flags.StringVar(&cfg.SMTPFrom, "smtp-from", "",
		"From address used in outgoing mail (required when smtp-addr is set).")
	flags.StringVar(&cfg.SMTPTo, "smtp-to", "",
		"Recipient address for contact submissions (required when smtp-addr is set).")
	flags.StringVar(&cfg.SubmissionsFile, "submissions-file", "submissions.jsonl",
		"Path to JSONL backup file for contact submissions. Empty disables file backup.")

	if err := flags.Parse(args[1:]); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func newServer(logger *slog.Logger, r *tmpl.Renderer, sink *submission.Sink) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, r, sink)
	return LoggerMiddleware(logger, mux)
}

func buildSink(cfg Config, logger *slog.Logger) (*submission.Sink, error) {
	var smtpCfg *submission.SMTPConfig
	if cfg.SMTPAddr != "" {
		if cfg.SMTPFrom == "" || cfg.SMTPTo == "" {
			return nil, fmt.Errorf("-smtp-addr is set but -smtp-from or -smtp-to is empty")
		}
		smtpCfg = &submission.SMTPConfig{
			Addr: cfg.SMTPAddr,
			User: cfg.SMTPUser,
			Pass: cfg.SMTPPass,
			From: cfg.SMTPFrom,
			To:   cfg.SMTPTo,
		}
	}
	return submission.New(smtpCfg, cfg.SubmissionsFile, logger), nil
}

func run(ctx context.Context, w io.Writer, e io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := newConfig(args)
	if err != nil {
		_, _ = fmt.Fprintf(e, "config error: %s\n", err)
		return err
	}

	logger := slog.New(slog.NewTextHandler(w, nil))

	renderer, err := tmpl.NewRenderer(os.DirFS("."))
	if err != nil {
		_, _ = fmt.Fprintf(e, "template registry error: %s\n", err)
		return err
	}

	sink, err := buildSink(cfg, logger)
	if err != nil {
		_, _ = fmt.Fprintf(e, "submission sink error: %s\n", err)
		return err
	}
	if !sink.Configured() {
		logger.Warn("no contact submission backend configured; submissions will only appear in the request log")
	} else {
		logger.Info("contact submission sink ready",
			slog.Bool("smtp", sink.SMTP != nil),
			slog.String("file", sink.FilePath))
	}

	srv := &http.Server{
		Addr:           net.JoinHostPort(cfg.Host, cfg.Port),
		Handler:        newServer(logger, renderer, sink),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		logger.Info("listening", slog.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			_, _ = fmt.Fprintf(e, "listen error: %s\n", err)
		}
	}()

	<-ctx.Done()
	logger.Info("shutting down")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		_, _ = fmt.Fprintf(e, "shutdown error: %s\n", err)
		return err
	}
	return nil
}

func main() {
	if err := run(context.Background(), os.Stdout, os.Stderr, os.Args); err != nil {
		os.Exit(1)
	}
}
