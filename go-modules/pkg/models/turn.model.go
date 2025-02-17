package models

import (
	"errors"
	"fmt"
)

type TurnPhase string

const (
	DrawCardsPhase  TurnPhase = "DRAW_CARDS_PHASE"
	PlaceCardsPhase TurnPhase = "PLACE_CARDS_PHASE"
	ActionPhase     TurnPhase = "ACTION_PHASE"
	EndPhase        TurnPhase = "END_PHASE"
)

type Turn struct {
	CurrentPlayer *Player
	Phase         TurnPhase
	PlayerIndex   int
}

func NewTurn(player *Player, playerIndex int) (*Turn, error) {
	if player == nil {
		return nil, errors.New("player must be provided")
	}

	if playerIndex < 0 || playerIndex > 1 {
		return nil, fmt.Errorf("invalid playerIndex %q: expected 0 or 1", playerIndex)
	}

	return &Turn{
		CurrentPlayer: player,
		Phase:         DrawCardsPhase,
		PlayerIndex:   playerIndex,
	}, nil
}

func (t *Turn) NextPhase() error {
	switch t.Phase {
	case DrawCardsPhase:
		t.Phase = PlaceCardsPhase
	case PlaceCardsPhase:
		t.Phase = ActionPhase
	case ActionPhase:
		t.Phase = EndPhase
	default:
		return fmt.Errorf("cannot advance to next phase because current phase is %s", t.Phase)
	}
	return nil
}
