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
		mux.Handle("/", loggingMiddleware(logger, printHeadersMiddleWare(logger, createRootHandler(logger, getRosterHandler, fileServer))))
	}
	mux.Handle("/roster", loggingMiddleware(logger, printHeadersMiddleWare(logger, createGetRosterHandler(logger))))
	mux.Handle("/roster/players-for-team", loggingMiddleware(logger, printHeadersMiddleWare(logger, createPlayersForTeamHandler(logger))))
	mux.Handle("/build-info", loggingMiddleware(logger, printHeadersMiddleWare(logger, createGetBuildInfoHandler(logger, buildInfo))))
}
