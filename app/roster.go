package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/StephenGriese/roster/handlers"
)

func Run(
	ctx context.Context,
	_ io.Reader,
	stdout, stderr io.Writer,
	getenv func(string) string,
	getwd func() (string, error),
) error {

	logger := slog.New(slog.NewJSONHandler(stdout, nil))
	wd, _ := getwd()
	logger.Info("Starting server", "working dir", wd)

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
		log.Printf("listening on %s\n", addr)
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
