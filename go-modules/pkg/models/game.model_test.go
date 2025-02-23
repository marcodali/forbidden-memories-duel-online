package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")

	deckA, _ := NewDeck(playerA, [40]*CardInstance{})
	deckB, _ := NewDeck(playerB, [40]*CardInstance{})

	game, err := NewGame([2]*Deck{deckA, deckB})
	assert.NoError(t, err)
	assert.NotNil(t, game)
	assert.NotNil(t, game.Board)
	assert.Equal(t, GameReadyToStart, game.State)
	assert.Equal(t, 0, game.CurrentTurn.PlayerIndex)
	assert.Equal(t, deckA, game.Decks[game.CurrentTurn.PlayerIndex])
	assert.Equal(t, playerA, game.CurrentTurn.CurrentPlayer)
	assert.Nil(t, game.eventChan)
	assert.Equal(t, time.Duration(0), game.DuelDuration)
}

func TestNewGameWithInvalidDecks(t *testing.T) {
	playerA, _ := NewPlayer("PlayerA")

	deckA, _ := NewDeck(playerA, [40]*CardInstance{})

	_, err := NewGame([2]*Deck{deckA, nil})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "both decks must be provided")
}

func TestStartGame(t *testing.T) {
	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")

	deckA, _ := NewDeck(playerA, [40]*CardInstance{})
	deckB, _ := NewDeck(playerB, [40]*CardInstance{})

	game, _ := NewGame([2]*Deck{deckA, deckB})

	assert.Equal(t, GameReadyToStart, game.State)
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

func TestFinishGame(t *testing.T) {
	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")

	deckA, _ := NewDeck(playerA, [40]*CardInstance{})
	deckB, _ := NewDeck(playerB, [40]*CardInstance{})

	game, _ := NewGame([2]*Deck{deckA, deckB})

	err := game.Finish()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "game cannot be finished in its current state")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)

	game.Start()
	err = game.Finish()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, game.DuelDuration, 1*time.Nanosecond, "duel duration is at least 1 nano second")
	_, successReadingEventFromChannel := <-game.eventChan
	assert.False(t, successReadingEventFromChannel, "eventChan should be closed")

	// finish the game twice causes an error too
	err = game.Finish()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "game cannot be finished in its current state")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestNextTurnSuccess(t *testing.T) {
	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")

	deckA, _ := NewDeck(playerA, [40]*CardInstance{})
	deckB, _ := NewDeck(playerB, [40]*CardInstance{})

	game, _ := NewGame([2]*Deck{deckA, deckB})

	// game always starts with playerA(0-index) as the current turn
	game.Start()

	// advance through all the phases of the current player's turn
	assert.Equal(t, DrawCardsPhase, game.CurrentTurn.Phase)
	game.CurrentTurn.NextPhase()
	assert.Equal(t, PlaceCardsPhase, game.CurrentTurn.Phase)
	game.CurrentTurn.NextPhase()
	assert.Equal(t, ActionPhase, game.CurrentTurn.Phase)
	game.CurrentTurn.NextPhase()
	assert.Equal(t, EndPhase, game.CurrentTurn.Phase)

	// call NextTurn() only after the player completes all the phases of their turn
	nextTurnDeck, err := game.NextTurn()
	assert.NoError(t, err)
	assert.Equal(t, 1, game.CurrentTurn.PlayerIndex)
	assert.Equal(t, nextTurnDeck, deckB)
	assert.Equal(t, playerB, game.CurrentTurn.CurrentPlayer)

	// The following are 3 different ways to achieve the same objective, choose wisely
	// Option (1) iterate using range with slice/array
	/* for range []string{"draw", "place", "action"} {
		game.CurrentTurn.NextPhase()
	} */

	// Option (2) you can count on me like 1, 2, 3
	/* for range 3 {
		game.CurrentTurn.NextPhase()
	} */

	// Option (3) Hardcode internal phase state
	game.CurrentTurn.Phase = EndPhase

	// when playerB completes their turn, the next turn should be playerA and so on
	nextTurnDeck, err = game.NextTurn()
	assert.NoError(t, err)
	assert.Equal(t, 0, game.CurrentTurn.PlayerIndex)
	assert.Equal(t, nextTurnDeck, deckA)
	assert.Equal(t, playerA, game.CurrentTurn.CurrentPlayer)
}

func TestNextTurnWithInvalidGameState(t *testing.T) {
	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")

	deckA, _ := NewDeck(playerA, [40]*CardInstance{})
	deckB, _ := NewDeck(playerB, [40]*CardInstance{})

	game, _ := NewGame([2]*Deck{deckA, deckB})

	_, err := game.NextTurn()
	assert.Error(t, err)

	// the error here is that game.Start() was never called
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot advance turn in the current game state")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestNextTurnWithInvalidPhaseState(t *testing.T) {
	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")

	deckA, _ := NewDeck(playerA, [40]*CardInstance{})
	deckB, _ := NewDeck(playerB, [40]*CardInstance{})

	game, _ := NewGame([2]*Deck{deckA, deckB})
	game.Start()

	// the error here is that the phase of the turn for playerA never advanced
	_, err := game.NextTurn()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot advance turn in the current turn phase")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestAddEvent(t *testing.T) {
	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")

	deckA, _ := NewDeck(playerA, [40]*CardInstance{})
	deckB, _ := NewDeck(playerB, [40]*CardInstance{})

	game, _ := NewGame([2]*Deck{deckA, deckB})

	// Trying to create a sample event, but fail
	event, err := NewEvent("TestEvent", map[string]any{"key": "value"})
	assert.Nil(t, event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid event type")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)

	// Again trying to create an event successfully
	event, err = NewEvent(EventDeckShuffled, map[string]any{"deck": deckA})
	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, event.Data["status"], SOEPristine)

	// Trying to add the event to the game, but failing
	err = game.AddEvent(event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("events can be added only during %s phase", GameInProgress))

	// Again trying to add the event to the game successfully
	game.Start()
	assert.NotNil(t, game.eventChan)

	go game.AddEvent(event)

	// wait for the event to be processed
	time.Sleep(1 * time.Millisecond)

	assert.Equal(t, event.Data["status"], SOECompleted)
}
