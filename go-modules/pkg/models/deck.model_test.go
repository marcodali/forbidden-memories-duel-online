package models

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var singletonForDeckModel sync.Once

func initializeDeckTestSuite() {
	singletonForDeckModel.Do(func() {
		_, filename, _, _ := runtime.Caller(0)
		fmt.Println("This setup code executes only one time for the file", filepath.Base(filename))
		CleanRegistry()
	})
}

func TestNewDeck(t *testing.T) {
	initializeDeckTestSuite()
	player, err := NewPlayer("TestPlayer")
	assert.NoError(t, err)
	assert.NotNil(t, player)
	cards := [40]*CardInstance{}
	for i := range 40 {
		cards[i] = &CardInstance{}
	}

	deck, err := NewDeck(player, cards)
	assert.NoError(t, err)
	assert.NotNil(t, deck)
	assert.Equal(t, 40, len(deck.RemainingCards))

	// valid deck type
	assert.Nil(t, deck.DeckType)
	deck.SetDeckType(DeckTypeYami)
	assert.Equal(t, DeckTypeYami, *deck.DeckType)
	assert.NotNil(t, deck.DeckType)

	// invalid deck type
	err = deck.SetDeckType("InvalidDeckType")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid deck type")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestNewDeckWithoutPlayer(t *testing.T) {
	initializeDeckTestSuite()
	cards := [40]*CardInstance{}
	_, err := NewDeck(nil, cards)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "player cannot be empty")
}

func TestMoveCardsFromRemainingToHand(t *testing.T) {
	initializeDeckTestSuite()
	player, err := NewPlayer("TestPlayer")
	assert.NoError(t, err)
	assert.NotNil(t, player)
	cards := [40]*CardInstance{}
	for i := range 40 {
		cards[i] = &CardInstance{}
	}

	deck, err := NewDeck(player, cards)
	assert.NoError(t, err)
	assert.NotNil(t, deck)
	assert.Equal(t, 40, len(deck.RemainingCards))

	err = deck.MoveCardsFromRemainingToHand(5)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(deck.HandCards))
	assert.Equal(t, 35, len(deck.RemainingCards))

	err = deck.MoveCardsFromRemainingToHand(-1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot move zero or less cards")

	err = deck.MoveCardsFromRemainingToHand(6)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot move more than 5 cards at a time")

	for range 35 {
		deck.MoveCardsFromRemainingToHand(1)
	}

	err = deck.MoveCardsFromRemainingToHand(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not enough remaining cards to move")
}
