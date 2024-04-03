package handlers

import (
	"log/slog"
	"net/http"
)

func HandleGetFavicon(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Getting favicon")
		http.ServeFile(w, r, "web/static/favicon.ico")
	}
}
