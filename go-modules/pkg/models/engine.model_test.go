package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEngine(t *testing.T) {
	engine := NewEngine()

	assert.NotNil(t, engine.activeGames)
	assert.Equal(t, 0, engine.totalGamesProcessed)
	assert.Greater(t, time.Now(), engine.startTime)
}

func TestAddGame(t *testing.T) {
	engine := NewEngine()

	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")
	deck1, _ := NewDeck(playerA, [40]*CardInstance{})
	deck2, _ := NewDeck(playerB, [40]*CardInstance{})
	game, _ := NewGame([2]*Deck{deck1, deck2})

	// fail adding a Game
	err := engine.AddGame(game)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("only games with State = %s can be added to the engine, got %s", GameInProgress, GameReadyToStart))

	// success adding a Game
	game.Start()
	err = engine.AddGame(game)
	assert.NoError(t, err)
	assert.Equal(t, 1, engine.GetActiveGamesCount(), "should be 1 active game")
}

func TestGetGame(t *testing.T) {
	engine := NewEngine()

	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")
	deck1, _ := NewDeck(playerA, [40]*CardInstance{})
	deck2, _ := NewDeck(playerB, [40]*CardInstance{})
	game, _ := NewGame([2]*Deck{deck1, deck2})

	// Add Game
	game.Start()
	engine.AddGame(game)

	// Get Game
	retrievedGame, err := engine.GetActiveGame(game.ID)
	assert.NoError(t, err)
	assert.Equal(t, game, retrievedGame)
	assert.EqualValues(t, game, retrievedGame)

	// Get non-existent Game
	_, err = engine.GetActiveGame("Not a valid game id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot get active game because not found")
}

func TestRemoveGame(t *testing.T) {
	engine := NewEngine()

	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")
	deck1, _ := NewDeck(playerA, [40]*CardInstance{})
	deck2, _ := NewDeck(playerB, [40]*CardInstance{})
	game, _ := NewGame([2]*Deck{deck1, deck2})

	// Add Game
	game.Start()
	engine.AddGame(game)

	// Remove non-existent Game
	err := engine.RemoveGame("Not a valid game id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot remove game because not found")

	// Remove a valid Game with invalid state
	err = engine.RemoveGame(game.ID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("only games with State = %s can be removed from the engine, got %s", GameFinished, GameInProgress))

	// Remove a valid Game with valid state
	game.Finish()
	err = engine.RemoveGame(game.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, engine.GetTotalGamesProcessed(), "should be 1 game processed")
}

func TestGetEngineUptime(t *testing.T) {
	engine := NewEngine()

	// Wait for a short duration
	time.Sleep(1 * time.Microsecond)

	assert.GreaterOrEqual(t, engine.GetEngineUptime(), time.Duration(1*time.Microsecond), "at least 1 microsecond should have elapsed since the engine started")
}
