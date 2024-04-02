package nhle_test

import (
	"testing"

	"github.com/StephenGriese/roster/nhle"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	ps := nhle.NewPlayerService()
	players, err := ps.Players()
	assert.NoError(t, err)
	assert.True(t, len(players) > 0)
}
