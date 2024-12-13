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

func Form() g.Node {
	return h.Form(
		g.Attr("hx-trigger", "change"),
		g.Attr("hx-get", "/roster/players-for-team"),
		g.Attr("hx-target", "#player-table-body"),
		g.Attr("hx-swap", "outerHTML"),
		TeamSelect(),
		SeasonSelect(),
		SortChoice(),
	)
}

func SortChoice() g.Node {
	return h.FieldSet(
		h.Legend(g.Text("Sort by")),
		h.Input(
			h.Type("radio"),
			h.ID("number"),
			h.Name("sort"),
			h.Value("number"),
			h.Checked(),
		),
		h.Label(h.For("number"), g.Text("Number")),
		h.Input(
			h.Type("radio"),
			h.ID("name"),
			h.Name("sort"),
			h.Value("name"),
		),
		h.Label(h.For("name"), g.Text("Name")),
		h.Input(
			h.Type("radio"),
			h.ID("position"),
			h.Name("sort"),
			h.Value("position"),
		),
		h.Label(h.For("position"), g.Text("Position")),
		h.Input(
			h.Type("radio"),
			h.ID("height"),
			h.Name("sort"),
			h.Value("height"),
		),
		h.Label(h.For("height"), g.Text("Height")),
		h.Input(
			h.Type("radio"),
			h.ID("weight"),
			h.Name("sort"),
			h.Value("weight"),
		),
		h.Label(h.For("weight"), g.Text("Weight")),
		h.Input(
			h.Type("radio"),
			h.ID("age"),
			h.Name("sort"),
			h.Value("age"),
		),
		h.Label(h.For("age"), g.Text("Age")),
	)
}

func SeasonSelect() g.Node {
	return h.Select(
		h.Name("season"),
		h.ID("season-select"),
		h.Option(h.Value(""), g.Text("Current")),
		h.Option(h.Value("20232024"), g.Text("2023-2024")),
		h.Option(h.Value("20222023"), g.Text("2022-2023")),
		h.Option(h.Value("20212022"), g.Text("2021-2022")),
		h.Option(h.Value("20202021"), g.Text("2020-2021")),
		h.Option(h.Value("20192020"), g.Text("2019-2020")),
		h.Option(h.Value("20182019"), g.Text("2018-2019")),
		h.Option(h.Value("20172018"), g.Text("2017-2018")),
		h.Option(h.Value("20162017"), g.Text("2016-2017")),
		h.Option(h.Value("20152016"), g.Text("2015-2016")),
		h.Option(h.Value("20142015"), g.Text("2014-2015")),
		h.Option(h.Value("20132014"), g.Text("2013-2014")),
		h.Option(h.Value("20122013"), g.Text("2012-2013")),
		h.Option(h.Value("20112012"), g.Text("2011-2012")),
		h.Option(h.Value("20102011"), g.Text("2010-2011")),
		h.Option(h.Value("20092010"), g.Text("2009-2010")),
		h.Option(h.Value("20082009"), g.Text("2008-2009")),
		h.Option(h.Value("20072008"), g.Text("2007-2008")),
		h.Option(h.Value("20062007"), g.Text("2006-2007")),
		h.Option(h.Value("20052006"), g.Text("2005-2006")),
		h.Option(h.Value("20042005"), g.Text("2004-2005")),
		h.Option(h.Value("20032004"), g.Text("2003-2004")),
		h.Option(h.Value("20022003"), g.Text("2002-2003")),
		h.Option(h.Value("20012002"), g.Text("2001-2002")),
		h.Option(h.Value("20002001"), g.Text("2000-2001")),
		h.Option(h.Value("19991900"), g.Text("1999-2000")),
		h.Option(h.Value("19981999"), g.Text("1998-1999")),
		h.Option(h.Value("19971998"), g.Text("1997-1998")),
		h.Option(h.Value("19961997"), g.Text("1996-1997")),
		h.Option(h.Value("19951996"), g.Text("1995-1996")),
		h.Option(h.Value("19941995"), g.Text("1994-1995")),
		h.Option(h.Value("19931994"), g.Text("1993-1994")),
		h.Option(h.Value("19921993"), g.Text("1992-1993")),
		h.Option(h.Value("19911992"), g.Text("1991-1992")),
		h.Option(h.Value("19901991"), g.Text("1990-1991")),
		h.Option(h.Value("19891990"), g.Text("1989-1990")),
		h.Option(h.Value("19881989"), g.Text("1988-1989")),
		h.Option(h.Value("19871988"), g.Text("1987-1988")),
		h.Option(h.Value("19861987"), g.Text("1986-1987")),
		h.Option(h.Value("19851986"), g.Text("1985-1986")),
		h.Option(h.Value("19841985"), g.Text("1984-1985")),
		h.Option(h.Value("19831984"), g.Text("1983-1984")),
		h.Option(h.Value("19821983"), g.Text("1982-1983")),
		h.Option(h.Value("19811982"), g.Text("1981-1982")),
		h.Option(h.Value("19801981"), g.Text("1980-1981")),
	)
}

func TeamSelect() g.Node {

	l := h.Label(
		g.Text("Team"),
		h.For("team-select"),
	)

	s := h.Select(
		h.Name("team"),
		h.ID("team-select"),
		MetropolitanOptGroup(),
		AtlanticOptGroup(),
		CentralOptGroup(),
		PacificOptGroup(),
	)
	return g.Group([]g.Node{l, s})
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
		h.ID("player-table-body"),
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
