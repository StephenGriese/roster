package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/StephenGriese/roster/handlers"
)

type BuildInfo struct {
	Builder   string
	BuildTime string
	Goversion string
	Version   string
}

func Run(
	ctx context.Context,
	_ io.Reader,
	stdout, stderr io.Writer,
	getenv func(string) string,
	getwd func() (string, error),
	buildInfo BuildInfo,
) error {

	logger := slog.New(slog.NewJSONHandler(stdout, nil))
	wd, _ := getwd()
	logger.Info("Starting server", "working dir", wd)
	logger.Info("Build info", "builder", buildInfo.Builder, "buildTime", buildInfo.BuildTime, "goversion", buildInfo.Goversion, "version", buildInfo.Version)

	addr := fmt.Sprintf(":%s", getenv("PORT"))
	if addr == ":" {
		return errors.New("missing PORT environment variable")
	}

	mux := http.NewServeMux()

	mux.Handle("/*", http.FileServer(http.Dir("./web/static/")))

	mux.HandleFunc("/roster", handlers.HandleGetRoster(logger))

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		logger.Info("starting http server", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			_, _ = fmt.Fprintf(stderr, "error shutting down http server: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		// make a new context for the Shutdown (thanks Alessandro Rosetti)
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			_, _ = fmt.Fprintf(stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
