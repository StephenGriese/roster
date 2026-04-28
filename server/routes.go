package server

import (
	"log/slog"
	"net/http"

	"github.com/StephenGriese/roster/nhle"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	buildInfo BuildInfo,
	ps nhle.PlayerService,
) {
	{
		getRosterHandler := createGetRosterHandler(logger, ps)
		fileServer := http.FileServer(http.Dir("./web/static/"))
		mux.Handle("/", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createRootHandler(logger, getRosterHandler, fileServer))))
	}
	mux.Handle("/roster", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createGetRosterHandler(logger, ps))))
	mux.Handle("/roster/players-for-team", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createPlayersForTeamHandler(logger, ps))))
	mux.Handle("/roster/download", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createDownloadRosterHandler(logger, ps))))
	mux.Handle("/player-search", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createPlayerSearchPageHandler(logger))))
	mux.Handle("/player-search/search", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createPlayerSearchHandler(logger, ps))))
	mux.Handle("/player-search/career", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createPlayerCareerHandler(logger, ps))))
	mux.Handle("/build-info", loggingMiddleware(logger, printXForwardedForMiddleWare(logger, createGetBuildInfoHandler(logger, buildInfo))))
}
