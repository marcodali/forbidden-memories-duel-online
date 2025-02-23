package models

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Only games with State = GameInProgress can live inside activeGames
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

// game should have started already to be added
func (e *Engine) AddGame(game *Game) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if game.State != GameInProgress {
		return fmt.Errorf("only games with State = %s can be added to the engine, got %s", GameInProgress, game.State)
	}
	e.activeGames[game.ID] = game
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

// game should be finished already to be removed
func (e *Engine) RemoveGame(gameID string) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	game, exists := e.activeGames[gameID]
	if !exists {
		return errors.New("cannot remove game because not found")
	}
	if game.State != GameFinished {
		return fmt.Errorf("only games with State = %s can be removed from the engine, got %s", GameFinished, game.State)
	}
	delete(e.activeGames, gameID)
	e.totalGamesProcessed++
	return nil
}
