package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventProhibitOpponentToAtackFnWithSuccess(t *testing.T) {
	playerA, _ := NewPlayer("PlayerA")
	playerB, _ := NewPlayer("PlayerB")

	event, _ := NewEvent(EventProhibitOpponentToAtack, map[string]any{
		"opponent": playerB,
		"turns":    3,
	})

	deckA, _ := NewDeck(playerA, [40]*CardInstance{})
	deckB, _ := NewDeck(playerB, [40]*CardInstance{})

	game, _ := NewGame([2]*Deck{deckA, deckB})
	game.Start()

	// playerA first turn and plays the card 348 - Swords of Revealing Light
	game.CurrentTurn.Phase = EndPhase
	go game.AddEvent(event)
	game.NextTurn()

	// wait for the event to be processed
	time.Sleep(1 * time.Millisecond)

	// playerB first turn
	assert.Equal(t, 3, game.CurrentTurn.CurrentPlayer.RemainingTurnsToAtack)
	game.CurrentTurn.Phase = EndPhase
	game.NextTurn()

	// playerA second turn
	game.CurrentTurn.Phase = EndPhase
	game.NextTurn()

	// playerB second turn
	assert.Equal(t, 2, game.CurrentTurn.CurrentPlayer.RemainingTurnsToAtack)
}

func TestInvalidEventProhibitOpponentToAtackFn(t *testing.T) {
	CleanRegistry()
	event, _ := NewEvent(EventProhibitOpponentToAtack, map[string]any{"key": "value"})

	// hardcode an invalid type
	event.Type = "Not a valid event type"
	err := EventProhibitOpponentToAtackFn(event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid event type")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)

	// hardcode an invalid status
	event.Data["status"] = SOEPristine
	event.Type = EventProhibitOpponentToAtack
	err = EventProhibitOpponentToAtackFn(event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid event status")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)

	// required properties check
	event.Data["status"] = SOEProcessing
	event.Type = EventProhibitOpponentToAtack
	err = EventProhibitOpponentToAtackFn(event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "turns missing")

	event.Data["turns"] = 5
	err = EventProhibitOpponentToAtackFn(event)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "opponent player missing")
}
