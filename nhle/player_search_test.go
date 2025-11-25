package nhle_test

import (
	"testing"

	"github.com/StephenGriese/roster/nhle"
	"github.com/stretchr/testify/assert"
)

func TestSearchPlayers(t *testing.T) {
	ps := nhle.NewPlayerService()
	results, err := ps.SearchPlayers("Crosby")
	assert.NoError(t, err)
	assert.True(t, len(results) > 0, "Expected to find at least one player named Crosby")

	// Check that we got Sidney Crosby
	found := false
	for _, p := range results {
		if p.Name == "Sidney Crosby" {
			found = true
			assert.Equal(t, "C", p.PositionCode)
			break
		}
	}
	assert.True(t, found, "Expected to find Sidney Crosby in search results")
}

func TestGetPlayerCareer(t *testing.T) {
	ps := nhle.NewPlayerService()
	// Sidney Crosby's player ID
	career, err := ps.GetPlayerCareer(8471675)
	assert.NoError(t, err)
	assert.NotNil(t, career)
	assert.Equal(t, "Sidney", career.FirstName)
	assert.Equal(t, "Crosby", career.LastName)
	assert.True(t, len(career.SeasonTotals) > 0, "Expected Crosby to have season stats")

	// Check that he has Pittsburgh in his team history
	teams := career.GetUniqueTeams()
	assert.True(t, len(teams) > 0, "Expected at least one team")
	found := false
	for _, team := range teams {
		if team == "Pittsburgh Penguins" {
			found = true
			break
		}
	}
	assert.True(t, found, "Expected Pittsburgh Penguins in career teams")
}
