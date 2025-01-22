package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"sync"
	"time"

	"github.com/StephenGriese/roster/roster"

	"github.com/StephenGriese/roster/nhle"
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
	stdout io.Writer,
	getenv func(string) string,
	getwd func() (string, error),
	buildInfo BuildInfo,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(stdout, nil))

	wd, _ := getwd()
	logger.Info("Starting server", "working dir", wd)
	logger.Info("Build info", "builder", buildInfo.Builder, "buildTime", buildInfo.BuildTime, "goversion", buildInfo.Goversion, "version", buildInfo.Version)

	addr := fmt.Sprintf(":%s", getenv("PORT"))
	if addr == ":" {
		return errors.New("missing PORT environment variable")
	}

	srv := NewServer(logger, buildInfo)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: srv,
	}

	go func() {
		logger.Info("starting http server", "addr", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Warn("error starting http server", "error", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		logger.Info("shutting down http server")
		// make a new context for the Shutdown (thanks Alessandro Rosetti)
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logger.Warn("error shutting down http server", "error", err)
		}
	}()
	wg.Wait()
	return nil
}

func NewServer(
	logger *slog.Logger,
	buildInfo BuildInfo,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		logger,
		buildInfo,
	)
	return mux
}

func createGetBuildInfoHandler(logger *slog.Logger, buildInfo BuildInfo) http.Handler {
	logger.Info("creating build info handler")
	return http.HandlerFunc( // This is a plain old Go type conversion. T(v) converts the value v to the type T.
		func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Getting build info")
			err := Page(BuildInfoContent(buildInfo)).Render(w)
			if err != nil {
				http.Error(w, "Error", http.StatusInternalServerError)
			}
		})
}

// createRootHandler solves a special problem. The "/" pattern matches everything, so we need to check that we're at the root here.
// The root handler will serve the roster page, while the file server will serve the static files.
func createRootHandler(logger *slog.Logger, getRosterHandler, fileServer http.Handler) http.Handler {
	logger.Info("creating root handler")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/":
				logger.Info("Root")
				getRosterHandler.ServeHTTP(w, r)
			default:
				logger.Info("FileServer")
				fileServer.ServeHTTP(w, r)
			}
		})
}

func createGetRosterHandler(logger *slog.Logger) http.Handler {
	logger.Info("creating get roster handler")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Getting roster")
			ps := nhle.NewPlayerService()
			players, err := ps.Players("PHI", "current")
			if err != nil {
				http.Error(w, "Error", http.StatusInternalServerError)
				return
			}
			sort.Slice(players, func(i, j int) bool {
				return players[i].SweaterNumber < players[j].SweaterNumber
			})
			err = Page(Form(), Table(players)).Render(w)
			if err != nil {
				logger.Error("Error rendering view", "error", err)
				http.Error(w, "Error", http.StatusInternalServerError)
			}
		})
}

func createPlayersForTeamHandler(logger *slog.Logger) http.Handler {
	logger.Info("creating players for team handler")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Info("playersForTeam")
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error", http.StatusInternalServerError)
				return
			}
			team := r.FormValue("team")
			sortBy := r.FormValue("sort")
			season := r.FormValue("season")
			logger.Info("formValues", "team", team, "sortBy", sortBy, "season", season)
			ps := nhle.NewPlayerService()
			players, err := ps.Players(team, season)
			if err != nil {
				http.Error(w, "Error", http.StatusInternalServerError)
				return
			}

			sort.Slice(players, makeSortFunc(players, sortBy))

			err = TableBody(players).Render(w)
			if err != nil {
				logger.Error("Error rendering view", "error", err)
				http.Error(w, "Error", http.StatusInternalServerError)
			}
		})
}

func makeSortFunc(players []roster.Player, sortBy string) func(i, j int) bool {
	switch sortBy {
	case "number":
		return func(i, j int) bool {
			return players[i].SweaterNumber < players[j].SweaterNumber
		}
	case "name":
		return func(i, j int) bool {
			return players[i].FullName() < players[j].FullName()
		}
	case "position":
		return func(i, j int) bool {
			if players[i].Position == players[j].Position {
				return players[i].SweaterNumber < players[j].SweaterNumber
			}
			return players[i].Position < players[j].Position
		}
	case "height":
		return func(i, j int) bool {
			if players[i].HeightInInches == players[j].HeightInInches {
				return players[i].SweaterNumber < players[j].SweaterNumber
			}
			return players[i].HeightInInches < players[j].HeightInInches
		}
	case "weight":
		return func(i, j int) bool {
			if players[i].WeightInPounds == players[j].WeightInPounds {
				return players[i].SweaterNumber < players[j].SweaterNumber
			}
			return players[i].WeightInPounds < players[j].WeightInPounds
		}
	case "age":
		return func(i, j int) bool {
			iAge := players[i].Age()
			jAge := players[j].Age()
			if iAge == jAge {
				return players[i].SweaterNumber < players[j].SweaterNumber
			}
			return iAge < jAge
		}
	default:
		return func(i, j int) bool {
			return players[i].SweaterNumber < players[j].SweaterNumber
		}
	}
}

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		logger.Info(
			"Got a request",
			"method", r.Method,
			"path", r.URL.Path,
			"remoteAddr", r.RemoteAddr)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func printHeadersMiddleWare(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, values := range r.Header {
			for _, value := range values {
				logger.Info("Header", "key", key, "value", value)
			}
		}
		next.ServeHTTP(w, r)
	})
}
