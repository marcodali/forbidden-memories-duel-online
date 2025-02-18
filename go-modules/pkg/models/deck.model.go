package models

import (
	"errors"
	"fmt"
	"slices"
)

// DeckType represents predefined archetypes for decks
type DeckType string

const (
	DeckTypeFemale    DeckType = "FEMALE"
	DeckTypeMountain  DeckType = "MOUNTAIN"
	DeckTypeYami      DeckType = "YAMI"
	DeckTypeForest    DeckType = "FOREST"
	DeckTypeAqua      DeckType = "AQUA"
	DeckTypeWarrior   DeckType = "WARRIOR"
	DeckTypeWasteland DeckType = "WASTELAND"
	DeckTypeGeneric   DeckType = "GENERIC"
)

var validDeckTypes = []DeckType{DeckTypeFemale, DeckTypeMountain, DeckTypeYami, DeckTypeForest, DeckTypeAqua, DeckTypeWarrior, DeckTypeWasteland, DeckTypeGeneric}

// Deck represents a collection of cards that a player can use in a game
type Deck struct {
	Player             *Player
	DeckType           *DeckType
	RemainingCards     []*CardInstance
	HandCards          []*CardInstance
	ActiveCardsOnBoard []*CardInstance
	DestroyedCards     []*CardInstance
}

// creates a new deck for a player with exactly 40 cards
func NewDeck(player *Player, cards [40]*CardInstance) (*Deck, error) {
	if player == nil {
		return nil, errors.New("player cannot be empty")
	}

	// Convert array to slice
	cardsSlice := cards[:]

	return &Deck{
		Player:             player,
		RemainingCards:     cardsSlice,
		HandCards:          []*CardInstance{},
		ActiveCardsOnBoard: []*CardInstance{},
		DestroyedCards:     []*CardInstance{},
	}, nil
}

// moves a specified number of cards from remaining cards to hand
func (d *Deck) MoveCardsFromRemainingToHand(count int) error {
	if count <= 0 {
		return errors.New("cannot move zero or less cards")
	}
	if count > 5 {
		return errors.New("cannot move more than 5 cards at a time")
	}
	if len(d.RemainingCards) < count {
		return errors.New("not enough remaining cards to move")
	}

	d.HandCards = append(d.HandCards, d.RemainingCards[:count]...)
	d.RemainingCards = d.RemainingCards[count:]
	return nil
}

func (d *Deck) SetDeckType(deckType DeckType) error {
	// verify if deckType is in the list of validDeckTypes
	if slices.Contains(validDeckTypes, deckType) {
		d.DeckType = &deckType
		return nil
	}
	return fmt.Errorf("invalid deck type %q: expected one of [%v]", deckType, validDeckTypes)
}
