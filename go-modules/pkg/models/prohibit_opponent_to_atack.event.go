package models

import (
	"errors"
	"fmt"
)

// The card 348 - Swords of Revealing Light trigger this event
func EventProhibitOpponentToAtackFn(event *Event) error {
	if event.Type != EventProhibitOpponentToAtack {
		return fmt.Errorf("invalid event type %s: expected %s", event.Type, EventProhibitOpponentToAtack)
	}
	if event.Data["status"] != SOEProcessing {
		return fmt.Errorf("invalid event status %s: expected %s", event.Data["status"], SOEProcessing)
	}

	// gathering requirements
	turns, turnsExists := event.Data["turns"].(int)
	if !turnsExists {
		return errors.New("turns missing")
	}
	opponent, playerExists := event.Data["opponent"].(*Player)
	if !playerExists {
		return errors.New("opponent player missing")
	}

	fmt.Println("Prohibiting the Opponent To Atack...")
	opponent.RemainingTurnsToAtack = turns
	return nil
}
