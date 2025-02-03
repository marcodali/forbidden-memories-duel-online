package models

import (
	"errors"
	"fmt"
	"time"
)

type GameState string

const (
	GameReadyToStart GameState = "READY_TO_START"
	GameInProgress   GameState = "IN_PROGRESS"
	GameCompleted    GameState = "COMPLETED"
)

type Game struct {
	Decks       [2]*Deck
	CurrentTurn int
	State       GameState
	StartTime   time.Time
}

func NewGame(decks [2]*Deck) (*Game, error) {
	if decks[0] == nil || decks[1] == nil {
		return nil, errors.New("both decks must be provided")
	}

	return &Game{
		Decks:       decks,
		CurrentTurn: 0, // represents the index of the current player
		State:       GameReadyToStart,
		StartTime:   time.Now(),
	}, nil
}

func (g *Game) Start() error {
	if g.State != GameReadyToStart {
		return fmt.Errorf("game cannot be started in its current state, expected: %s, got: %s", GameReadyToStart, g.State)
	}

	g.State = GameInProgress
	g.StartTime = time.Now()
	return nil
}

func (g *Game) NextTurn() (*Deck, error) {
	if g.State != GameInProgress {
		return nil, fmt.Errorf("cannot advance turn in the current game state, expected: %s, got: %s", GameInProgress, g.State)
	}

	// Alternate between players and return the current player's deck
	g.CurrentTurn = (g.CurrentTurn + 1) % 2
	return g.Decks[g.CurrentTurn], nil
}
