package models

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var singletonForBoardModel sync.Once

func initializeBoardTestSuite() {
	singletonForBoardModel.Do(func() {
		_, filename, _, _ := runtime.Caller(0)
		fmt.Println("This setup code executes only one time for the file", filepath.Base(filename))
		CleanRegistry()
	})
}

var PLAYER_A = 0
var PLAYER_B = 1

func LoadReal722CardsFromYAML() {
	if GetCardRegistry().GetCard(722) != nil {
		return // cards were already loaded
	}
	data, _ := os.ReadFile("../utils/cards.yaml")
	GetCardRegistry().LoadCardsfromYAML(data)
}

func TestSetCardAtIndexPositionOnlyMonsterTypes(t *testing.T) {
	initializeBoardTestSuite()
	LoadReal722CardsFromYAML()
	t.Run("should place card in monster zone correctly", func(t *testing.T) {
		board := NewBoard()

		turnOfPlayerA := PLAYER_A
		indexPosition := 3
		babyDragon, err := NewCardInstance(4)
		assert.NoError(t, err)
		// Test valid position
		cardStatePlayerA := &CardState{Card: babyDragon, FaceUp: true, IndexPosition: indexPosition}
		err = board.SetCardAtIndexPosition(cardStatePlayerA, turnOfPlayerA)
		assert.NoError(t, err)
		assert.Nil(t, board.MonsterZones[turnOfPlayerA][indexPosition-1])
		assert.NotNil(t, board.MonsterZones[turnOfPlayerA][indexPosition])
		assert.Equal(t, cardStatePlayerA, board.MonsterZones[turnOfPlayerA][indexPosition])
		assert.Nil(t, board.MonsterZones[turnOfPlayerA][indexPosition+1])
		// Test invalid position
		cardStatePlayerA = &CardState{Card: babyDragon, FaceUp: true, IndexPosition: 5}
		err = board.SetCardAtIndexPosition(cardStatePlayerA, turnOfPlayerA)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid card index position")

		// to see this error message, run the test with -v flag
		t.Logf("Error: %v", err)

		turnOfPlayerB := PLAYER_B
		indexPosition = 1
		tyhone, err := NewCardInstance(13)
		assert.NoError(t, err)
		// Test valid position
		cardStatePlayerB := &CardState{Card: tyhone, FaceUp: false, IndexPosition: indexPosition}
		err = board.SetCardAtIndexPosition(cardStatePlayerB, turnOfPlayerB)
		assert.NoError(t, err)
		assert.Nil(t, board.MonsterZones[turnOfPlayerB][indexPosition-1])
		assert.NotNil(t, board.MonsterZones[turnOfPlayerB][indexPosition])
		assert.Equal(t, cardStatePlayerB, board.MonsterZones[turnOfPlayerB][indexPosition])
		assert.Nil(t, board.MonsterZones[turnOfPlayerB][indexPosition+1])
		// Test invalid position
		cardStatePlayerB = &CardState{Card: tyhone, FaceUp: false, IndexPosition: -1}
		err = board.SetCardAtIndexPosition(cardStatePlayerB, turnOfPlayerA)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid card index position")

		// to see this error message, run the test with -v flag
		t.Logf("Error: %v", err)
	})
}

func TestSetCardAtIndexPositionOnlyMagicTypes(t *testing.T) {
	initializeBoardTestSuite()
	LoadReal722CardsFromYAML()
	board := NewBoard()
	turnOfPlayerA := PLAYER_A
	turnOfPlayerB := PLAYER_B
	hamburgerRecipe, _ := NewCardInstance(677)
	hornOfLight, _ := NewCardInstance(313)
	badReactionToSimochi, _ := NewCardInstance(688)
	shadowSpell, _ := NewCardInstance(669)

	t.Run("should place a ritual card in magic zone correctly", func(t *testing.T) {
		indexPosition := 0
		cardStatePlayerA := &CardState{Card: hamburgerRecipe, FaceUp: true, IndexPosition: indexPosition}
		err := board.SetCardAtIndexPosition(cardStatePlayerA, turnOfPlayerA)
		assert.NoError(t, err)
		assert.Equal(t, cardStatePlayerA, board.MagicTrapZones[turnOfPlayerA][indexPosition])
	})

	t.Run("should place an equip card in magic zone correctly", func(t *testing.T) {
		indexPosition := 0
		cardStatePlayerB := &CardState{Card: hornOfLight, FaceUp: true, IndexPosition: indexPosition}
		err := board.SetCardAtIndexPosition(cardStatePlayerB, turnOfPlayerB)
		assert.NoError(t, err)
		assert.Equal(t, cardStatePlayerB, board.MagicTrapZones[turnOfPlayerB][indexPosition])
	})

	t.Run("should place a trap card in magic zone correctly", func(t *testing.T) {
		indexPosition := 4
		cardStatePlayerA := &CardState{Card: badReactionToSimochi, FaceUp: false, IndexPosition: indexPosition}
		err := board.SetCardAtIndexPosition(cardStatePlayerA, turnOfPlayerA)
		assert.NoError(t, err)
		assert.Equal(t, cardStatePlayerA, board.MagicTrapZones[turnOfPlayerA][indexPosition])
	})

	t.Run("should place a magic card in magic zone correctly", func(t *testing.T) {
		indexPosition := 4
		cardStatePlayerB := &CardState{Card: shadowSpell, FaceUp: false, IndexPosition: indexPosition}
		err := board.SetCardAtIndexPosition(cardStatePlayerB, turnOfPlayerB)
		assert.NoError(t, err)
		assert.Equal(t, cardStatePlayerB, board.MagicTrapZones[turnOfPlayerB][indexPosition])
	})

	t.Run("should fail at placing the card because invalid type", func(t *testing.T) {
		invalidCard, _ := NewCardInstance(1)
		invalidCard.Template.Type = "Not a valid type"
		cardStatePlayerA := &CardState{Card: invalidCard, FaceUp: false, IndexPosition: 3}
		err := board.SetCardAtIndexPosition(cardStatePlayerA, turnOfPlayerA)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid card type")

		// to see this error message, run the test with -v flag
		t.Logf("Error: %v", err)
	})
}
