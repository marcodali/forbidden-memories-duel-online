package models

import (
	"runtime"
	"fmt"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var singletonForEngineModel sync.Once

func initializeEngineTestSuite() {
	singletonForEngineModel.Do(func() {
		_, filename, _, _ := runtime.Caller(0)
		fmt.Println("This setup code executes only one time for the file", filepath.Base(filename))
		CleanRegistry()
	})
}

func TestNewEngine(t *testing.T) {
	initializeEngineTestSuite()
	engine := NewEngine()
	assert.NotNil(t, engine)
	assert.NotNil(t, engine.activeGames)
	assert.False(t, engine.startTime.IsZero())
}

func TestEngineStartAndShutdown(t *testing.T) {
	initializeEngineTestSuite()
	engine := NewEngine()

	err := engine.Start()
	assert.NoError(t, err)

	err = engine.Shutdown()
	assert.NoError(t, err)
	assert.Nil(t, engine.activeGames)
}

func TestCreateAndEndGame(t *testing.T) {
	initializeEngineTestSuite()
	engine := NewEngine()
	
	// Create players and decks
	player1 := createTestPlayer(t, "Player1")
	player2 := createTestPlayer(t, "Player2")
	deck1 := createTestDeck(t, player1)
	deck2 := createTestDeck(t, player2)

	// Create game
	game, err := engine.CreateGame(deck1, deck2)
	assert.NoError(t, err)
	assert.NotNil(t, game)
	assert.Equal(t, 1, engine.GetActiveGamesCount())

	// End game
	err = engine.EndGame(game.ID)
	assert.NoError(t, err)
	assert.Equal(t, 0, engine.GetActiveGamesCount())
	assert.Equal(t, int64(1), engine.GetTotalGamesProcessed())
}

func TestGetGame(t *testing.T) {
	initializeEngineTestSuite()
	engine := NewEngine()
	
	// Create test game
	player1 := createTestPlayer(t, "Player1")
	player2 := createTestPlayer(t, "Player2")
	deck1 := createTestDeck(t, player1)
	deck2 := createTestDeck(t, player2)
	game, _ := engine.CreateGame(deck1, deck2)

	// Test getting existing game
	retrieved, err := engine.GetGame(game.ID)
	assert.NoError(t, err)
	assert.NotNil(t, retrieved)
	assert.Equal(t, game, retrieved)

	// Test getting non-existent game
	retrieved, err = engine.GetGame("non-existent")
	assert.NoError(t, err)
	assert.Nil(t, retrieved)
}

func TestEngineMetrics(t *testing.T) {
	initializeEngineTestSuite()
	engine := NewEngine()

	// Create and finish some games
	for i := 0; i < 3; i++ {
		player1 := createTestPlayer(t, "Player1")
		player2 := createTestPlayer(t, "Player2")
		deck1 := createTestDeck(t, player1)
		deck2 := createTestDeck(t, player2)
		game, _ := engine.CreateGame(deck1, deck2)
		
		// Start the game to allow finishing
		err := game.Start()
		assert.NoError(t, err)
		
		time.Sleep(100 * time.Millisecond) // Simulate game duration
		err = engine.EndGame(game.ID)
		assert.NoError(t, err)
	}

	// Test metrics
	assert.Equal(t, 0, engine.GetActiveGamesCount())
	assert.Equal(t, int64(3), engine.GetTotalGamesProcessed())
	assert.True(t, engine.GetEngineUptime() > 0)
	assert.True(t, engine.GetAverageDuelDuration() > 0)

	stateCount := engine.GetGamesByState()
	assert.NotNil(t, stateCount)
}

// Helper function to create a test player
func createTestPlayer(t *testing.T, username string) *Player {
	player, err := NewPlayer(username)
	assert.NoError(t, err)
	return player
}

// Helper function to create a test deck
func createTestDeck(t *testing.T, player *Player) *Deck {
	InitializeCardRegistryWithFakeYAMLData()
	cards := [40]*CardInstance{}
	for i := 0; i < 40; i++ {
		card, err := NewCardInstance(1001)
		assert.NoError(t, err)
		cards[i] = card
	}
	deck, err := NewDeck(player, cards)
	assert.NoError(t, err)
	return deck
}
