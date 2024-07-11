package handlers

import (
	"github.com/StephenGriese/roster/roster"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"log/slog"
	"net/http"
	"sort"
	"strconv"

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
		err = Page(Table(players)).Render(w)
		if err != nil {
			logger.Error("Error rendering view", "error", err)
			http.Error(w, "Error", http.StatusInternalServerError)
		}
	}
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
