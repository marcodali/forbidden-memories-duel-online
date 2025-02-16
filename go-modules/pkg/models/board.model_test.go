package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var PLAYER_1 = 0
var PLAYER_2 = 1

func LoadReal722CardsFromYAML() {
	data, _ := os.ReadFile("../utils/cards.yaml")
	GetCardRegistry().LoadCardsfromYAML(data)
}

func TestSetCardAtIndexPositionOnlyMonsterTypes(t *testing.T) {
	LoadReal722CardsFromYAML()
	t.Run("should place card in monster zone correctly", func(t *testing.T) {
		board := NewBoard()

		turnOfPlayer1 := PLAYER_1
		indexPosition := 3
		babyDragon, err := NewCardInstance(4)
		assert.NoError(t, err)
		// Test valid position
		cardStatePlayer1 := &CardState{Card: babyDragon, FaceUp: true, IndexPosition: indexPosition}
		err = board.SetCardAtIndexPosition(cardStatePlayer1, turnOfPlayer1)
		assert.NoError(t, err)
		assert.Nil(t, board.MonsterZones[turnOfPlayer1][indexPosition-1])
		assert.NotNil(t, board.MonsterZones[turnOfPlayer1][indexPosition])
		assert.Equal(t, cardStatePlayer1, board.MonsterZones[turnOfPlayer1][indexPosition])
		assert.Nil(t, board.MonsterZones[turnOfPlayer1][indexPosition+1])
		// Test invalid position
		cardStatePlayer1 = &CardState{Card: babyDragon, FaceUp: true, IndexPosition: 5}
		err = board.SetCardAtIndexPosition(cardStatePlayer1, turnOfPlayer1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid index position")

		// to see this error message, run the test with -v flag
		t.Logf("Error: %v", err)

		turnOfPlayer2 := PLAYER_2
		indexPosition = 1
		tyhone, err := NewCardInstance(13)
		assert.NoError(t, err)
		// Test valid position
		cardStatePlayer2 := &CardState{Card: tyhone, FaceUp: false, IndexPosition: indexPosition}
		err = board.SetCardAtIndexPosition(cardStatePlayer2, turnOfPlayer2)
		assert.NoError(t, err)
		assert.Nil(t, board.MonsterZones[turnOfPlayer2][indexPosition-1])
		assert.NotNil(t, board.MonsterZones[turnOfPlayer2][indexPosition])
		assert.Equal(t, cardStatePlayer2, board.MonsterZones[turnOfPlayer2][indexPosition])
		assert.Nil(t, board.MonsterZones[turnOfPlayer2][indexPosition+1])
		// Test invalid position
		cardStatePlayer2 = &CardState{Card: tyhone, FaceUp: false, IndexPosition: -1}
		err = board.SetCardAtIndexPosition(cardStatePlayer2, turnOfPlayer1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid index position")

		// to see this error message, run the test with -v flag
		t.Logf("Error: %v", err)
	})
}

func TestSetCardAtIndexPositionOnlyMagicTypes(t *testing.T) {
	LoadReal722CardsFromYAML()
	board := NewBoard()
	turnOfPlayer1 := PLAYER_1
	turnOfPlayer2 := PLAYER_2
	hamburgerRecipe, _ := NewCardInstance(677)
	hornOfLight, _ := NewCardInstance(313)
	badReactionToSimochi, _ := NewCardInstance(688)
	shadowSpell, _ := NewCardInstance(669)

	t.Run("should place a ritual card in magic zone correctly", func(t *testing.T) {
		indexPosition := 0
		cardStatePlayer1 := &CardState{Card: hamburgerRecipe, FaceUp: true, IndexPosition: indexPosition}
		err := board.SetCardAtIndexPosition(cardStatePlayer1, turnOfPlayer1)
		assert.NoError(t, err)
		assert.Equal(t, cardStatePlayer1, board.MagicTrapZones[turnOfPlayer1][indexPosition])
	})

	t.Run("should place an equip card in magic zone correctly", func(t *testing.T) {
		indexPosition := 0
		cardStatePlayer2 := &CardState{Card: hornOfLight, FaceUp: true, IndexPosition: indexPosition}
		err := board.SetCardAtIndexPosition(cardStatePlayer2, turnOfPlayer2)
		assert.NoError(t, err)
		assert.Equal(t, cardStatePlayer2, board.MagicTrapZones[turnOfPlayer2][indexPosition])
	})

	t.Run("should place a trap card in magic zone correctly", func(t *testing.T) {
		indexPosition := 4
		cardStatePlayer1 := &CardState{Card: badReactionToSimochi, FaceUp: false, IndexPosition: indexPosition}
		err := board.SetCardAtIndexPosition(cardStatePlayer1, turnOfPlayer1)
		assert.NoError(t, err)
		assert.Equal(t, cardStatePlayer1, board.MagicTrapZones[turnOfPlayer1][indexPosition])
	})

	t.Run("should place a magic card in magic zone correctly", func(t *testing.T) {
		indexPosition := 4
		cardStatePlayer2 := &CardState{Card: shadowSpell, FaceUp: false, IndexPosition: indexPosition}
		err := board.SetCardAtIndexPosition(cardStatePlayer2, turnOfPlayer2)
		assert.NoError(t, err)
		assert.Equal(t, cardStatePlayer2, board.MagicTrapZones[turnOfPlayer2][indexPosition])
	})

	t.Run("should fail at placing the card because invalid type", func(t *testing.T) {
		invalidCard, _ := NewCardInstance(1)
		invalidCard.Template.Type = "Not a valid type"
		cardStatePlayer1 := &CardState{Card: invalidCard, FaceUp: false, IndexPosition: 3}
		err := board.SetCardAtIndexPosition(cardStatePlayer1, turnOfPlayer1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid card type")

		// to see this error message, run the test with -v flag
		t.Logf("Error: %v", err)
	})
}
