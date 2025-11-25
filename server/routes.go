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
		mux.Handle("/", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createRootHandler(logger, getRosterHandler, fileServer))))
	}
	mux.Handle("/roster", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createGetRosterHandler(logger))))
	mux.Handle("/roster/players-for-team", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createPlayersForTeamHandler(logger))))
	mux.Handle("/player-search", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createPlayerSearchPageHandler(logger))))
	mux.Handle("/player-search/search", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createPlayerSearchHandler(logger))))
	mux.Handle("/player-search/career", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createPlayerCareerHandler(logger))))
	mux.Handle("/build-info", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createGetBuildInfoHandler(logger, buildInfo))))
}
