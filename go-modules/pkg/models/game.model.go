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
	GameFinished     GameState = "FINISHED"
)

type Game struct {
	Decks        [2]*Deck
	Board        *Board
	CurrentTurn  *Turn
	State        GameState
	StartTime    time.Time
	DuelDuration time.Duration
	eventChan    chan *Event
}

func NewGame(decks [2]*Deck) (*Game, error) {
	if decks[0] == nil || decks[1] == nil {
		return nil, errors.New("both decks must be provided")
	}

	decks[0].Player.LifePoints = 8000
	decks[1].Player.LifePoints = 8000

	turn, _ := NewTurn(decks[0].Player, 0)
	game := &Game{
		Decks:       decks,
		Board:       NewBoard(),
		CurrentTurn: turn,
		State:       GameReadyToStart,
		StartTime:   time.Now(),
	}

	return game, nil
}

func (g *Game) Start() error {
	if g.State != GameReadyToStart {
		return fmt.Errorf("game cannot be started in its current state, expected: %s, got: %s", GameReadyToStart, g.State)
	}

	g.State = GameInProgress
	g.StartTime = time.Now()
	g.eventChan = make(chan *Event)

	// Launch the event processing goroutine
	go g.processEvents()

	return nil
}

func (g *Game) Finish() error {
	if g.State != GameInProgress {
		return fmt.Errorf("game cannot be finished in its current state, expected: %s, got: %s", GameInProgress, g.State)
	}

	g.State = GameFinished
	g.DuelDuration = time.Since(g.StartTime)
	close(g.eventChan)
	return nil
}

// calling normal AddEvent(e) could return errors and will block the code execution until the event is consumed
// in the other hand, calling go AddEvent(e) as a gorutine will execute the code in an async(non-blocking) way
// on the background which sounds great but errors cannot be catched anymore.
func (g *Game) AddEvent(event *Event) error {
	// events can only be added after GameReadyToStart phase and prior to GameFinished phase
	if g.State != GameInProgress {
		return fmt.Errorf("events can be added only during %s phase", GameInProgress)
	}
	event.Data["status"] = SOEEnqueued
	g.eventChan <- event
	return nil
}

// this is a forever loop running in the background if executed as gorutine
func (g *Game) processEvents() {
	for event := range g.eventChan {
		if processingEventFunction, functionExists := validEventTypes[event.Type]; functionExists {
			event.Data["status"] = SOEProcessing
			processingEventFunction(event)
			event.Data["status"] = SOECompleted
		}
	}
}

func (g *Game) NextTurn() (*Deck, error) {
	if g.State != GameInProgress {
		return nil, fmt.Errorf("cannot advance turn in the current game state, expected: %s, got: %s", GameInProgress, g.State)
	}

	if g.CurrentTurn.Phase != EndPhase {
		return nil, fmt.Errorf("cannot advance turn in the current turn phase, expected: %s, got: %s", EndPhase, g.CurrentTurn.Phase)
	}

	if g.CurrentTurn.CurrentPlayer.RemainingTurnsToAtack > 0 {
		g.CurrentTurn.CurrentPlayer.RemainingTurnsToAtack -= 1
	}

	nextPlayerIndex := (g.CurrentTurn.PlayerIndex + 1) % 2
	nextPlayer := g.Decks[nextPlayerIndex].Player
	g.CurrentTurn, _ = NewTurn(nextPlayer, nextPlayerIndex)

	return g.Decks[nextPlayerIndex], nil
}
