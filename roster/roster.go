package roster

import (
	"fmt"
	"time"
)

type Position uint

const (
	Forward Position = iota
	Defense
	Goalie
)

func (p Position) String() string {
	switch p {
	case Forward:
		return "forward"
	case Defense:
		return "defense"
	case Goalie:
		return "goalie"
	default:
		return "unknown"
	}
}

type Player struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	SweaterNumber  int       `json:"sweaterNumber"`
	Position       Position  `json:"position"`
	HeightInInches int       `json:"heightInInches"`
	WeightInPounds int       `json:"weightInPounds"`
	BirthDate      time.Time `json:"birthDate"`
}

func (p Player) HeightInFeetAndInches() (feet, inches int) {
	feet = p.HeightInInches / 12
	inches = p.HeightInInches % 12
	return
}

func FeetAndInchesToString(feet, inches int) string {
	return fmt.Sprintf("%d'%d\"", feet, inches)
}

func (p Player) Age() int {
	now := time.Now()
	years := now.Year() - p.BirthDate.Year()
	if now.YearDay() < p.BirthDate.YearDay() {
		years--
	}
	return years
}

func (p Player) FullName() string {
	return p.LastName + ", " + p.FirstName
}

// PlayerSearchResult represents a player from search results
type PlayerSearchResult struct {
	PlayerID       int    `json:"playerId"`
	Name           string `json:"name"`
	TeamAbbrev     string `json:"teamAbbrev"`
	PositionCode   string `json:"positionCode"`
	Active         bool   `json:"active"`
	LastTeamAbbrev string `json:"lastTeamAbbrev"`
}

// PlayerCareer represents a player's career information
type PlayerCareer struct {
	PlayerID          int           `json:"playerId"`
	FirstName         string        `json:"firstName"`
	LastName          string        `json:"lastName"`
	SweaterNumber     int           `json:"sweaterNumber"`
	Position          string        `json:"position"`
	HeadShot          string        `json:"headshot"`
	TeamLogo          string        `json:"teamLogo"`
	CareerGamesPlayed int           `json:"careerGamesPlayed"`
	SeasonTotals      []SeasonStats `json:"seasonTotals"`
}

// SeasonStats represents stats for a single season
type SeasonStats struct {
	Season              int     `json:"season"`
	TeamName            string  `json:"teamName"`
	LeagueAbbrev        string  `json:"leagueAbbrev"`
	GamesPlayed         int     `json:"gamesPlayed"`
	Goals               int     `json:"goals"`
	Assists             int     `json:"assists"`
	Points              int     `json:"points"`
	PlusMinus           int     `json:"plusMinus"`
	Wins                int     `json:"wins"`
	Losses              int     `json:"losses"`
	GoalsAgainstAverage float64 `json:"goalsAgainstAvg"`
	SavePctg            float64 `json:"savePctg"`
}

func (pc PlayerCareer) FullName() string {
	return pc.FirstName + " " + pc.LastName
}

// GetUniqueTeams returns a list of unique teams the player has played for in the NHL
func (pc PlayerCareer) GetUniqueTeams() []string {
	teamMap := make(map[string]bool)
	teams := make([]string, 0)

	for _, season := range pc.SeasonTotals {
		if season.LeagueAbbrev == "NHL" {
			if !teamMap[season.TeamName] {
				teamMap[season.TeamName] = true
				teams = append(teams, season.TeamName)
			}
		}
	}

	return teams
}
