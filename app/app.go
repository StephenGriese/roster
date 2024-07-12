package app

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
	"strconv"
	"sync"
	"time"

	"github.com/StephenGriese/roster/nhle"
	"github.com/StephenGriese/roster/roster"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
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

	mux := http.NewServeMux()
	addRoutes(mux, logger, buildInfo)
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		logger.Info("starting http server", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Warn("error shutting down http server", "error", err)
		}
	}()
	wg.Wait()
	return nil
}

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	buildInfo BuildInfo,
) {
	mux.Handle("/roster", createGetRosterHandler(logger))
	mux.Handle("/build-info", createGetBuildInfoHandler(logger, buildInfo))
	mux.Handle("/*", http.FileServer(http.Dir("./web/static/")))
}

func createGetBuildInfoHandler(logger *slog.Logger, buildInfo BuildInfo) http.Handler {
	logger.Info("creating build info handler")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Getting build info")
			err := Page(BuildInfoContent(buildInfo)).Render(w)
			if err != nil {
				http.Error(w, "Error", http.StatusInternalServerError)
			}
		})
}

func createGetRosterHandler(logger *slog.Logger) http.Handler {
	type thing struct {
		counter int
	}
	theThing := thing{}
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			theThing.counter++
			logger.Info("Getting roster", "counter", theThing.counter)
			ps := nhle.NewPlayerService()
			players, err := ps.Players()
			if err != nil {
				http.Error(w, "Error", http.StatusInternalServerError)
				return
			}
			sort.Slice(players, func(i, j int) bool {
				return players[i].SweaterNumber < players[j].SweaterNumber
			})
			err = Page(Table(players)).Render(w)
			if err != nil {
				logger.Error("Error rendering view", "error", err)
				http.Error(w, "Error", http.StatusInternalServerError)
			}
		})
}

func Page(body g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    "Roster",
		Language: "en",
		Head: []g.Node{
			h.Script(h.Src("/js/htmx-1.9.11.js")),
		},
		Body: []g.Node{
			Container(
				Prose(body)),
		},
	})
}

func BuildInfoContent(info BuildInfo) g.Node {
	return h.Dl(
		h.Dt(g.Text("Builder")),
		h.Dd(g.Text(info.Builder)),
		h.Dt(g.Text("BuildTime")),
		h.Dd(g.Text(info.BuildTime)),
		h.Dt(g.Text("Goversion")),
		h.Dd(g.Text(info.Goversion)),
		h.Dt(g.Text("Version")),
		h.Dd(g.Text(info.Version)),
	)
}

func Table(players []roster.Player) g.Node {
	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Number")),
				h.Th(g.Text("LastName")),
				h.Th(g.Text("FirstName")),
				h.Th(g.Text("Position")),
			),
		),
		h.TBody(
			g.Group(g.Map(players, func(p roster.Player) g.Node {
				return h.Tr(
					h.Td(g.Text(strconv.Itoa(p.SweaterNumber))),
					h.Td(g.Text(p.LastName)),
					h.Td(g.Text(p.FirstName)),
					h.Td(g.Text(p.Position.String())))
			})),
		),
	)
}

func Container(children ...g.Node) g.Node {
	return h.Div(g.Group(children))
}

func Prose(children ...g.Node) g.Node {
	return h.Div(g.Group(children))
}
