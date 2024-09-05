package main

import (
	"github.com/StephenGriese/roster/roster"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"strconv"
)

func Page(body ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    "Roster",
		Language: "en",
		Head: []g.Node{
			h.Script(h.Src("/js/htmx-1.9.11.js")),
		},
		Body: body,
	})
}

func Table(players []roster.Player) g.Node {
	return h.Table(
		h.ID("player-table"),
		h.THead(
			h.Tr(
				h.Th(g.Text("Number")),
				h.Th(g.Text("LastName")),
				h.Th(g.Text("FirstName")),
				h.Th(g.Text("Position")),
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
	return h.Select(
		g.Attr("hx-get", "/roster/players-for-team"),
		g.Attr("hx-target", "#player-table"),
		h.Option(h.Value("flyers"), g.Text("flyers")),
		h.Option(h.Value("penguins"), g.Text("penguins")),
	)
}
