package app

import (
	"log/slog"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	buildInfo BuildInfo,
) {
	mux.Handle("/roster", createGetRosterHandler(logger))
	mux.Handle("/roster/players-for-team", createPlayersForTeamHandler(logger))
	mux.Handle("/build-info", createGetBuildInfoHandler(logger, buildInfo))
	mux.Handle("/", http.FileServer(http.Dir("./web/static/")))
}
