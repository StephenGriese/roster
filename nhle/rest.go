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
	BaseURLV1 = "https://api-web.nhle.com/v1/roster/%s/current"
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

func (ps PlayerService) Players(team string) ([]roster.Player, error) {
	s := fmt.Sprintf(BaseURLV1, team)
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
