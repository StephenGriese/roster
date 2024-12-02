package server

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
			h.Link(h.Rel("stylesheet"), h.Href("https://cdn.simplecss.org/simple.min.css")),
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
				h.Th(h.Style("width: 10%"), g.Text("Number")),
				h.Th(h.Style("width: 35%"), g.Text("Name")),
				h.Th(h.Style("width: 15%"), g.Text("Position")),
				h.Th(h.Style("width: 10%"), g.Text("Height")),
				h.Th(h.Style("width: 10%"), g.Text("Weight")),
				h.Th(h.Style("width: 10%"), g.Text("Age")),
			),
		),
		TableBody(players),
	)
}

func TableBody(players []roster.Player) g.Node {
	return h.TBody(
		g.Group(g.Map(players, func(p roster.Player) g.Node {
			s := roster.FeetAndInchesToString(p.HeightInFeetAndInches())
			return h.Tr(
				h.Td(g.Text(strconv.Itoa(p.SweaterNumber))),
				h.Td(g.Text(p.FullName())),
				h.Td(g.Text(p.Position.String())),
				h.Td(g.Text(s)),
				h.Td(g.Text(strconv.Itoa(p.WeightInPounds))),
				h.Td(g.Text(strconv.Itoa(p.Age()))),
			)
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
		MetropolitanOptGroup(),
		AtlanticOptGroup(),
		CentralOptGroup(),
		PacificOptGroup(),
	)
	return g.Group([]g.Node{l, br, s})
}

func AtlanticOptGroup() g.Node {
	return h.OptGroup(
		h.LabelAttr("Atlantic"),
		h.Option(h.Value("BOS"), g.Text("bruins")),
		h.Option(h.Value("BUF"), g.Text("sabres")),
		h.Option(h.Value("DET"), g.Text("red wings")),
		h.Option(h.Value("FLA"), g.Text("panthers")),
		h.Option(h.Value("MTL"), g.Text("canadiens")),
		h.Option(h.Value("OTT"), g.Text("senators")),
		h.Option(h.Value("TBL"), g.Text("lightning")),
		h.Option(h.Value("TOR"), g.Text("maple leafs")),
	)
}

func MetropolitanOptGroup() g.Node {
	return h.OptGroup(
		h.LabelAttr("Metropolitan"),
		h.Option(h.Value("CAR"), g.Text("hurricanes")),
		h.Option(h.Value("CBJ"), g.Text("blue jackets")),
		h.Option(h.Value("NJD"), g.Text("devils")),
		h.Option(h.Value("NYI"), g.Text("islanders")),
		h.Option(h.Value("NYR"), g.Text("rangers")),
		h.Option(h.Value("PHI"), g.Text("flyers"), h.Selected()),
		h.Option(h.Value("PIT"), g.Text("penguins")),
		h.Option(h.Value("WSH"), g.Text("capitals")),
	)
}

func CentralOptGroup() g.Node {
	return h.OptGroup(
		h.LabelAttr("Central"),
		h.Option(h.Value("CHI"), g.Text("blackhawks")),
		h.Option(h.Value("COL"), g.Text("avalanche")),
		h.Option(h.Value("DAL"), g.Text("stars")),
		h.Option(h.Value("MIN"), g.Text("wild")),
		h.Option(h.Value("NSH"), g.Text("predators")),
		h.Option(h.Value("STL"), g.Text("blues")),
		h.Option(h.Value("UTA"), g.Text("utah hockey club")),
		h.Option(h.Value("WPG"), g.Text("jets")),
	)
}

func PacificOptGroup() g.Node {
	return h.OptGroup(
		h.LabelAttr("Pacific"),
		h.Option(h.Value("ANA"), g.Text("ducks")),
		h.Option(h.Value("CGY"), g.Text("flames")),
		h.Option(h.Value("EDM"), g.Text("oilers")),
		h.Option(h.Value("LAK"), g.Text("kings")),
		h.Option(h.Value("SJS"), g.Text("sharks")),
		h.Option(h.Value("SEA"), g.Text("kraken")),
		h.Option(h.Value("VAN"), g.Text("canucks")),
		h.Option(h.Value("VGK"), g.Text("golden knights")),
	)
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
