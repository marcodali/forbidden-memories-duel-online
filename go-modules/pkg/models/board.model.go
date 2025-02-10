package models

import (
	"fmt"
)

// represents different card zones on the board
type Zone string

const (
	ZoneMonster   Zone = "MONSTER"
	ZoneMagicTrap Zone = "MAGIC_TRAP_EQUIP"
	ZoneField     Zone = "BATTLE_FIELD"
)

type CardState struct {
	Card          *CardInstance
	FaceUp        bool
	IndexPosition int
}

// represents the complete playing field for both players
type Board struct {
	MonsterZones   [2][5]*CardState // [playerTurn][position]
	MagicTrapZones [2][5]*CardState // [playerTurn][position]
	FieldZone      [2]*CardState    // current battle field monsters for both players
}

func NewBoard() *Board {
	return &Board{}
}

// places a card in its corresponding zone based on its type
func (b *Board) SetCardAtIndexPosition(state *CardState, currentTurn int) error {
	if state.IndexPosition < 0 || state.IndexPosition >= 5 {
		return fmt.Errorf("invalid index position: %d", state.IndexPosition)
	}

	switch state.Card.Template.Type {
	case TypeMagic, TypeTrap, TypeEquip, TypeRitual:
		b.MagicTrapZones[currentTurn][state.IndexPosition] = state
		return nil
	}

	if validMonsterTypes[state.Card.Template.Type] {
		b.MonsterZones[currentTurn][state.IndexPosition] = state
		return nil
	}

	return fmt.Errorf("invalid card type")
}
