package app

import (
	"os"
	"strconv"

	"github.com/StephenGriese/roster/roster"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

func Page(nodes ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    "Roster",
		Language: "en",
		Head: []g.Node{
			h.Script(h.Src("/js/htmx-1.9.11.js")),
			h.Link(h.Rel("stylesheet"), h.Href("https://cdn.simplecss.org/simple.css")),
		},
		Body: nodes,
	})
}

func Table(players []roster.Player) g.Node {
	return h.Table(
		h.Style("width: 100%"),
		h.ID("player-table"),
		h.THead(
			h.Tr(
				h.Th(h.Style("width: 25%"), g.Text("Number")),
				h.Th(h.Style("width: 25%"), g.Text("LastName")),
				h.Th(h.Style("width: 25%"), g.Text("FirstName")),
				h.Th(h.Style("width: 25%"), g.Text("Position")),
			),
		),
		TableBody(players),
	)
}

func TableBody(players []roster.Player) g.Node {
	return h.TBody(
		g.Group(g.Map(players, func(p roster.Player) g.Node {
			return h.Tr(
				h.Td(g.Text(strconv.Itoa(p.SweaterNumber))),
				h.Td(g.Text(p.LastName)),
				h.Td(g.Text(p.FirstName)),
				h.Td(g.Text(p.Position.String())))
		})),
	)
}

func TeamSelect() g.Node {

	l := h.Label(
		g.Text("Team"),
		h.For("team-select"),
	)

	br := h.Br()

	s := h.Select(
		g.Attr("hx-get", "/roster/players-for-team"),
		g.Attr("hx-target", "#player-table"),
		h.Name("team"),
		h.ID("team-select"),
		h.Option(h.Value("PHI"), g.Text("flyers")),
		h.Option(h.Value("PIT"), g.Text("penguins")),
	)
	return g.Group([]g.Node{l, br, s})
}

func BuildInfoContent(info BuildInfo) g.Node {

	wd, _ := os.Getwd()
	return h.Dl(
		h.Dt(g.Text(wd)),
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
