// Package nhle implements a PlayerService that uses api-web.nhle.com as a datasource
package nhle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/StephenGriese/roster/roster"
)

const (
	BaseURLV1       = "https://api-web.nhle.com/v1/roster/%s"
	SearchURLV1     = "https://search.d3.nhle.com/api/v1/search/player"
	PlayerLandingV1 = "https://api-web.nhle.com/v1/player/%d/landing"
)

func NewPlayerService() PlayerService {
	return PlayerService{
		baseURL: BaseURLV1,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type PlayerService struct {
	baseURL    string
	httpClient *http.Client
}

func (ps PlayerService) Players(team, season string) ([]roster.Player, error) {
	s := fmt.Sprintf(BaseURLV1, team)
	switch season {
	case "":
		s += "/current"
	case "current":
		s += "/current"
	default:
		s += fmt.Sprintf("/%s", season)
	}
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	res, err := ps.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	type Name struct {
		Default string `json:"default"`
	}

	type Player struct {
		ID             int        `json:"id"`
		FirstName      Name       `json:"firstName"`
		LastName       Name       `json:"lastName"`
		SweaterNumber  int        `json:"sweaterNumber"`
		HeightInInches int        `json:"heightInInches"`
		WeightInPounds int        `json:"weightInPounds"`
		BirthDate      CustomTime `json:"birthDate"`
	}

	toRosterPlayer := func(p Player, position roster.Position) roster.Player {
		return roster.Player{
			ID:             p.ID,
			FirstName:      p.FirstName.Default,
			LastName:       p.LastName.Default,
			SweaterNumber:  p.SweaterNumber,
			Position:       position,
			HeightInInches: p.HeightInInches,
			WeightInPounds: p.WeightInPounds,
			BirthDate:      p.BirthDate.Time,
		}
	}

	apiResp := struct {
		Forwards   []Player `json:"forwards"`
		Defensemen []Player `json:"defensemen"`
		Goalies    []Player `json:"goalies"`
	}{}

	if err = json.NewDecoder(res.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	numPlayers := len(apiResp.Forwards) + len(apiResp.Defensemen) + len(apiResp.Goalies)
	result := make([]roster.Player, 0, numPlayers)

	for _, p := range apiResp.Forwards {
		result = append(result, toRosterPlayer(p, roster.Forward))
	}
	for _, p := range apiResp.Defensemen {
		result = append(result, toRosterPlayer(p, roster.Defense))
	}
	for _, p := range apiResp.Goalies {
		result = append(result, toRosterPlayer(p, roster.Goalie))
	}
	return result, nil
}

// SearchPlayers searches for players by name
func (ps PlayerService) SearchPlayers(query string) ([]roster.PlayerSearchResult, error) {
	req, err := http.NewRequest("GET", SearchURLV1, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	q := req.URL.Query()
	q.Add("culture", "en-us")
	q.Add("limit", "20")
	q.Add("q", query)
	req.URL.RawQuery = q.Encode()

	res, err := ps.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	var apiResp []struct {
		PlayerID       string `json:"playerId"`
		Name           string `json:"name"`
		TeamAbbrev     string `json:"teamAbbrev"`
		PositionCode   string `json:"positionCode"`
		Active         bool   `json:"active"`
		LastTeamAbbrev string `json:"lastTeamAbbrev"`
	}

	if err = json.NewDecoder(res.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	result := make([]roster.PlayerSearchResult, 0, len(apiResp))
	for _, p := range apiResp {
		// Convert string PlayerID to int
		var playerID int
		if _, err := fmt.Sscanf(p.PlayerID, "%d", &playerID); err != nil {
			continue // Skip invalid player IDs
		}

		result = append(result, roster.PlayerSearchResult{
			PlayerID:       playerID,
			Name:           p.Name,
			TeamAbbrev:     p.TeamAbbrev,
			PositionCode:   p.PositionCode,
			Active:         p.Active,
			LastTeamAbbrev: p.LastTeamAbbrev,
		})
	}

	return result, nil
}

// GetPlayerCareer gets career information for a specific player
func (ps PlayerService) GetPlayerCareer(playerID int) (*roster.PlayerCareer, error) {
	url := fmt.Sprintf(PlayerLandingV1, playerID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := ps.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	type Name struct {
		Default string `json:"default"`
	}

	var apiResp struct {
		PlayerID      int    `json:"playerId"`
		FirstName     Name   `json:"firstName"`
		LastName      Name   `json:"lastName"`
		SweaterNumber int    `json:"sweaterNumber"`
		Position      string `json:"position"`
		HeadShot      string `json:"headshot"`
		TeamLogo      string `json:"teamLogo"`
		FeaturedStats struct {
			Season        int `json:"season"`
			RegularSeason struct {
				SubSeason struct {
					GamesPlayed int `json:"gamesPlayed"`
				} `json:"subSeason"`
				Career struct {
					GamesPlayed int `json:"gamesPlayed"`
				} `json:"career"`
			} `json:"regularSeason"`
		} `json:"featuredStats"`
		CareerTotals struct {
			RegularSeason struct {
				GamesPlayed int `json:"gamesPlayed"`
			} `json:"regularSeason"`
		} `json:"careerTotals"`
		SeasonTotals []struct {
			Season              int     `json:"season"`
			TeamName            Name    `json:"teamName"`
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
		} `json:"seasonTotals"`
	}

	if err = json.NewDecoder(res.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	seasons := make([]roster.SeasonStats, 0, len(apiResp.SeasonTotals))
	for _, s := range apiResp.SeasonTotals {
		seasons = append(seasons, roster.SeasonStats{
			Season:              s.Season,
			TeamName:            s.TeamName.Default,
			LeagueAbbrev:        s.LeagueAbbrev,
			GamesPlayed:         s.GamesPlayed,
			Goals:               s.Goals,
			Assists:             s.Assists,
			Points:              s.Points,
			PlusMinus:           s.PlusMinus,
			Wins:                s.Wins,
			Losses:              s.Losses,
			GoalsAgainstAverage: s.GoalsAgainstAverage,
			SavePctg:            s.SavePctg,
		})
	}

	return &roster.PlayerCareer{
		PlayerID:          apiResp.PlayerID,
		FirstName:         apiResp.FirstName.Default,
		LastName:          apiResp.LastName.Default,
		SweaterNumber:     apiResp.SweaterNumber,
		Position:          apiResp.Position,
		HeadShot:          apiResp.HeadShot,
		TeamLogo:          apiResp.TeamLogo,
		CareerGamesPlayed: apiResp.CareerTotals.RegularSeason.GamesPlayed,
		SeasonTotals:      seasons,
	}, nil
}
