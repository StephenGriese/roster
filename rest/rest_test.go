package rest_test

import (
	"github.com/StephenGriese/roster/rest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUsers(t *testing.T) {

	ps := rest.NewPlayerService()
	players, err := ps.Players()
	assert.NoError(t, err)
	assert.Equal(t, 10, len(players))
}
