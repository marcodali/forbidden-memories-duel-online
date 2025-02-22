package models

import (
	"errors"
	"sync"
	"time"
)

// Engine represents the game engine that manages active duels
type Engine struct {
	activeGames         map[string]*Game
	mutex               sync.RWMutex
	startTime           time.Time
	totalGamesProcessed int
}

func NewEngine() *Engine {
	return &Engine{
		activeGames: make(map[string]*Game),
		startTime:   time.Now(),
	}
}

// This is a write lock operation
func (e *Engine) AddGame(game *Game) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.activeGames[game.ID] = game
}

// This is a write lock operation
func (e *Engine) RemoveGame(gameID string) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	game, gameExists := e.activeGames[gameID]
	if !gameExists {
		return errors.New("cannot remove game because not found")
	}

	delete(e.activeGames, game.ID)
	e.totalGamesProcessed++
	return nil
}

func (e *Engine) GetActiveGame(gameID string) (*Game, error) {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	game, gameExists := e.activeGames[gameID]
	if !gameExists {
		return nil, errors.New("cannot get active game because not found")
	}

	return game, nil
}

func (e *Engine) GetActiveGamesCount() int {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return len(e.activeGames)
}

func (e *Engine) GetAverageDuelDuration() time.Duration {
	return time.Duration(0)
}

func (e *Engine) GetEngineUptime() time.Duration {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return time.Since(e.startTime)
}

func (e *Engine) GetTotalGamesProcessed() int {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.totalGamesProcessed
}
