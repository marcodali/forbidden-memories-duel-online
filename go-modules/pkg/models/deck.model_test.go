package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	player, err := NewPlayer("TestPlayer")
	assert.NoError(t, err)
	assert.NotNil(t, player)
	cards := [40]*CardInstance{}
	for i := 0; i < 40; i++ {
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
	assert.Contains(t, err.Error(), "invalid deck type")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestNewDeckInvalid(t *testing.T) {
	cards := [40]*CardInstance{}
	_, err := NewDeck(nil, cards)
	assert.Contains(t, err.Error(), "player cannot be empty")
}

func TestMoveCardsFromRemainingToHand(t *testing.T) {
	player, err := NewPlayer("TestPlayer")
	assert.NoError(t, err)
	assert.NotNil(t, player)
	cards := [40]*CardInstance{}
	for i := 0; i < 40; i++ {
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
	assert.Contains(t, err.Error(), "cannot move zero or less cards")

	err = deck.MoveCardsFromRemainingToHand(6)
	assert.Contains(t, err.Error(), "cannot move more than 5 cards at a time")

	for i := 0; i < 35; i++ {
		deck.MoveCardsFromRemainingToHand(1)
	}

	err = deck.MoveCardsFromRemainingToHand(1)
	assert.Contains(t, err.Error(), "not enough remaining cards to move")
}
