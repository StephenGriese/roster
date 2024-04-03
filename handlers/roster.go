package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"sort"

	"github.com/StephenGriese/roster/nhle"
	"github.com/StephenGriese/roster/views"
)

func HandleGetRoster(logger *slog.Logger, view *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Getting roster")
		ps := nhle.NewPlayerService()
		players, err := ps.Players()
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
		sort.Slice(players, func(i, j int) bool {
			return players[i].SweaterNumber < players[j].SweaterNumber
		})
		err = view.Render(w, players)
		if err != nil {
			logger.Error("Error rendering view", "error", err)
			http.Error(w, "Error", http.StatusInternalServerError)
		}
	}
}

func RosterHandler(w http.ResponseWriter, r *http.Request) {
	ps := nhle.NewPlayerService()
	players, err := ps.Players()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sort.Slice(players, func(i, j int) bool {
		return players[i].SweaterNumber < players[j].SweaterNumber
	})
	for _, b := range players {
		fmt.Fprintf(w, "%2d   %-25s %s\n", b.SweaterNumber, b.FirstName+" "+b.LastName, b.Position)
	}
}
