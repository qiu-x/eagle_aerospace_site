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

	"eagle_site_v2/tmpl"
)

type Config struct {
	Host string
	Port string
}

func newConfig(args []string) (Config, error) {
	var cfg Config
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.StringVar(&cfg.Host, "host", "0.0.0.0", "Bind address")
	flags.StringVar(&cfg.Port, "port", "8080", "Listen port")
	if err := flags.Parse(args[1:]); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func newServer(logger *slog.Logger, r *tmpl.Renderer) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, r)
	return LoggerMiddleware(logger, mux)
}

func run(ctx context.Context, w io.Writer, e io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := newConfig(args)
	if err != nil {
		fmt.Fprintf(e, "config error: %s\n", err)
		return err
	}

	logger := slog.New(slog.NewTextHandler(w, nil))

	renderer, err := tmpl.NewRenderer(os.DirFS("."))
	if err != nil {
		fmt.Fprintf(e, "template registry error: %s\n", err)
		return err
	}

	srv := &http.Server{
		Addr:           net.JoinHostPort(cfg.Host, cfg.Port),
		Handler:        newServer(logger, renderer),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		logger.Info("listening", slog.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintf(e, "listen error: %s\n", err)
		}
	}()

	<-ctx.Done()
	logger.Info("shutting down")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Fprintf(e, "shutdown error: %s\n", err)
		return err
	}
	return nil
}

func main() {
	if err := run(context.Background(), os.Stdout, os.Stderr, os.Args); err != nil {
		os.Exit(1)
	}
}
