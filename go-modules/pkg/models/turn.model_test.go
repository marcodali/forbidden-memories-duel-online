package models

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var singletonForTurnModel sync.Once

func initializeTurnTestSuite() {
	singletonForTurnModel.Do(func() {
		_, filename, _, _ := runtime.Caller(0)
		fmt.Println("This setup code executes only one time for the file", filepath.Base(filename))
		CleanRegistry()
	})
}

func TestNewTurn(t *testing.T) {
	initializeTurnTestSuite()

	player, _ := NewPlayer("TestPlayer")
	turn, err := NewTurn(player, 0)

	assert.NoError(t, err)
	assert.NotNil(t, turn)
	assert.Equal(t, player, turn.CurrentPlayer)
	assert.Equal(t, DrawCardsPhase, turn.Phase)
	assert.Equal(t, 0, turn.PlayerIndex)

	_, err = NewTurn(nil, 2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "player must be provided")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)

	_, err = NewTurn(player, -1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid playerIndex")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestNextPhaseWithInvalidState(t *testing.T) {
	initializeTurnTestSuite()

	player, _ := NewPlayer("TestPlayer")
	turn, _ := NewTurn(player, 0)

	turn.Phase = EndPhase
	err := turn.NextPhase()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot advance to next phase")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)

	turn.Phase = "Not a valid phase"
	err = turn.NextPhase()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot advance to next phase")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}
