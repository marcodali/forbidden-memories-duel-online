package models

import (
	"errors"
	"fmt"
)

func EventDeckShuffledFn(event *Event) error {
	if event.Type != EventDeckShuffled {
		return fmt.Errorf("invalid event type %s: expected %s", event.Type, EventDeckShuffled)
	}
	if event.Data["status"] != SOEProcessing {
		return fmt.Errorf("invalid event status %s: expected %s", event.Data["status"], SOEProcessing)
	}

	// gathering requirements
	deck, deckExists := event.Data["deck"].(*Deck)
	if !deckExists {
		return errors.New("deck missing")
	}

	fmt.Println("Shuffling cards on deck...", deck)
	return nil
}
