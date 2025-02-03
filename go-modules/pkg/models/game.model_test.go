package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	player1, _ := NewPlayer("Player1")
	player2, _ := NewPlayer("Player2")

	deck1, _ := NewDeck(player1, [40]*CardInstance{})
	deck2, _ := NewDeck(player2, [40]*CardInstance{})

	game, err := NewGame([2]*Deck{deck1, deck2})
	assert.NoError(t, err)
	assert.NotNil(t, game)
	assert.Equal(t, GameReadyToStart, game.State)
	assert.Equal(t, 0, game.CurrentTurn)
	assert.Equal(t, deck1, game.Decks[game.CurrentTurn])
	assert.Equal(t, player1, game.Decks[game.CurrentTurn].Player)
}

func TestNewGameWithInvalidDecks(t *testing.T) {
	player1, _ := NewPlayer("Player1")

	deck1, _ := NewDeck(player1, [40]*CardInstance{})

	_, err := NewGame([2]*Deck{deck1, nil})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "both decks must be provided")
}

func TestStartGame(t *testing.T) {
	player1, _ := NewPlayer("Player1")
	player2, _ := NewPlayer("Player2")

	deck1, _ := NewDeck(player1, [40]*CardInstance{})
	deck2, _ := NewDeck(player2, [40]*CardInstance{})

	game, _ := NewGame([2]*Deck{deck1, deck2})

	err := game.Start()
	assert.NoError(t, err)
	assert.Equal(t, GameInProgress, game.State)

	// Start the game again and ensure it fails
	err = game.Start()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "game cannot be started in its current state")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestNextTurn(t *testing.T) {
	player1, _ := NewPlayer("Player1")
	player2, _ := NewPlayer("Player2")

	deck1, _ := NewDeck(player1, [40]*CardInstance{})
	deck2, _ := NewDeck(player2, [40]*CardInstance{})

	game, _ := NewGame([2]*Deck{deck1, deck2})

	// game always starts with player1(0-index) as the current turn
	game.Start()

	// NextTurn() should be called when each player completes their turn
	nextTurnDeck, err := game.NextTurn()
	assert.NoError(t, err)
	assert.Equal(t, 1, game.CurrentTurn)
	assert.Equal(t, nextTurnDeck, deck2)

	// when player2 completes their turn, the next turn should be player1 and so on
	nextTurnDeck, err = game.NextTurn()
	assert.NoError(t, err)
	assert.Equal(t, 0, game.CurrentTurn)
	assert.Equal(t, nextTurnDeck, deck1)
}

func TestNextTurnInvalidState(t *testing.T) {
	player1, _ := NewPlayer("Player1")
	player2, _ := NewPlayer("Player2")

	deck1, _ := NewDeck(player1, [40]*CardInstance{})
	deck2, _ := NewDeck(player2, [40]*CardInstance{})

	game, _ := NewGame([2]*Deck{deck1, deck2})

	_, err := game.NextTurn()
	assert.Error(t, err)

	// the error here is that game.Start() was never called
	assert.Contains(t, err.Error(), "cannot advance turn in the current game state")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}
