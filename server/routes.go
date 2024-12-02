package server

import (
	"log/slog"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	buildInfo BuildInfo,
) {
	{
		getRosterHandler := createGetRosterHandler(logger)
		fileServer := http.FileServer(http.Dir("./web/static/"))
		mux.Handle("/", createRootHandler(logger, getRosterHandler, fileServer))
	}
	mux.Handle("/roster", createGetRosterHandler(logger))
	mux.Handle("/roster/players-for-team", createPlayersForTeamHandler(logger))
	mux.Handle("/build-info", createGetBuildInfoHandler(logger, buildInfo))
}
